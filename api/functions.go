package api

import (
	"encoding/json"
	"io"
	"os"
)

func LoadAnnonces() ([]Annonce, error) {
	annonces := []Annonce{}
	file, err := os.Open("./annonces.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &annonces)

	if err != nil {
		return nil, err
	}
	return annonces, nil
}

func LoadJsonAnnonces() ([]byte, error) {

	file, err := os.Open("./annonces.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func Contains[T comparable](arr []T, c T) bool {
	for _, item := range arr {
		if item == c {
			return true
		}
	}
	return false
}
