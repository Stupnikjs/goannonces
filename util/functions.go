package util

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Stupnikjs/zik/api"
)

func IsInSlice[T comparable](str T, arr []T) bool {
	for _, s := range arr {
		if str == s {
			return true
		}

	}
	return false
}

func CleanAllTempDir() error {

	curr, err := os.Getwd()

	if err != nil {
		return err
	}

	entries, err := os.ReadDir(path.Join(curr, "/static/download"))

	if err != nil {
		return err
	}
	for _, e := range entries {
		if strings.Contains(e.Name(), "temp") {
			err = os.RemoveAll(path.Join(curr, "/static/download", e.Name()))
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func ProcessMp3Name(mp3name string) string {
	firstSplit := strings.Split(mp3name, ".mp3")

	if len(firstSplit) == 1 {
		return mp3name
	}

	nodotmp3 := firstSplit[0]

	secsplit := strings.Split(nodotmp3, "/")
	fmt.Println(secsplit[len(secsplit)-1])
	return secsplit[len(secsplit)-1]

}

func GetJsonStructFromReq(r *http.Request) (*api.JsonReq, error) {

	return nil, nil
}
