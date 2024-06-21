package util

import (
	"os"
	"path"
	"strings"
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
