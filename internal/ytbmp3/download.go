package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func Download(videoID string, fileName string, client youtube.Client) error {
	video, err := client.GetVideo(videoID)
	if err != nil {
		return err
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	audioformat := len(formats) - 1
	stream, _, err := client.GetStream(video, &formats[audioformat])

	if err != nil {
		fmt.Println(err)
	}
	defer stream.Close()

	file, err := os.Create(fileName + ".mp3")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}

	return nil

}
