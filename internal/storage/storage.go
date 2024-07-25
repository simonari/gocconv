package storage

import (
	"fmt"
	"log"
	"os"
	"path"
)

type RatesFile struct {
	os.File
	path   string
	Stored uint8
}

func OpenRatesFile(filePath string) *RatesFile {
	dir, _ := path.Split(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)

		fmt.Printf("[+] Created directory at: %s\n", dir)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	info, err := file.Stat()

	if err != nil {
		log.Fatalf("Error getting stat %e", err)
	}

	if info.Size() == 0 {
		file.Write([]byte("{}"))
	}

	return &RatesFile{
		File: *file,
		path: filePath,
	}
}
