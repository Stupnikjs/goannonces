package ytbmp3

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func Download(videoID string) (string, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(videoID)
	if err != nil {
		return "", err
	}
	title := video.Title

	mp4Formats, err := getFormatFromVid(*video)

	if err != nil {
		return "", err
	}
	// Select the best quality MP4 format (you can define your own criteria for "best")
	bestFormat := selectBestFormat(mp4Formats)

	stream, _, err := client.GetStream(video, &bestFormat)

	if err != nil {
		return "", err
	}
	defer stream.Close()

	curr, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// check if temp dir exist

	temp, err := os.CreateTemp(path.Join(curr, "/static/download"), "tempvideofile")
	if err != nil {
		return "", err
	}

	defer temp.Close()

	_, err = io.Copy(temp, stream)
	if err != nil {
		return "", err
	}
	mp3name := title + ".mp3"

	fmt.Println(path.Join("/static", mp3name))
	if err = convertToMP3(temp.Name(), path.Join(curr, "/static/download", mp3name)); err != nil {
		return "", err
	}

	return mp3name, nil

}

func getFormatFromVid(video youtube.Video) ([]youtube.Format, error) {
	var mp4Formats []youtube.Format
	for _, format := range video.Formats.WithAudioChannels() {
		if strings.Contains(format.MimeType, "video/mp4") {
			mp4Formats = append(mp4Formats, format)
		}
	}
	if len(mp4Formats) == 0 {
		return nil, fmt.Errorf("no mp4 formats with audio available")
	}
	return mp4Formats, nil
}

func selectBestFormat(formats []youtube.Format) youtube.Format {
	// Select the format with the highest resolution
	bestFormat := formats[0]
	for _, format := range formats {
		if format.Height > bestFormat.Height {
			bestFormat = format
		}
	}
	return bestFormat
}

func convertToMP3(inputFile, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:a", "0", "-map", "a", outputFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg command failed: %w", err)
	}
	return nil
}
