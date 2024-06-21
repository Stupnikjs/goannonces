package util

import (
	"os"
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

func CleanAllTempDir(dirName string) error {

	entries, err := os.ReadDir(dirName)

	if err != nil {
		return err
	}
	for _, e := range entries {
		if strings.Contains(e.Name(), "temp") {
			err = os.RemoveAll(e.Name())
			if err != nil {
				return err
			}

		}
	}
	return nil
}
