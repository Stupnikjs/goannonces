package api

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Stupnikjs/goannonces/database"
)

type Application struct {
	Port int
	DB   database.DBRepo
}

func LoadAnnonces() ([]database.Annonce, error) {
	annonces := []database.Annonce{}
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
