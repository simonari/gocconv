package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Configuration struct {
	RatesStoragePath string `json:"RatesStoragePath"`
}

func Read() *Configuration {
	path := "config.json"

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("[!] Error: %s", err)
	}

	c := Configuration{}

	content, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("[!] Error: %s", err)
	}

	json.Unmarshal(content, &c)

	return &c
}
