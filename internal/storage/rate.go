package storage

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
	"vsimonari/gocconv/internal/core"
)

type RatesFile struct {
	os.File
	path   string
	Stored uint8
}

func OpenRatesFile(name string) *RatesFile {
	dir, _ := path.Split(name)

	os.Mkdir(dir, os.ModePerm)

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatalf("Error getting stat %e", err)
	}

	if fileInfo.Size() == 0 {
		writeDefaultData(file)
	}

	return &RatesFile{
		File: *file,
		path: name,
	}
}

func writeDefaultData(file *os.File) {
	data := initDefaultData()

	_, err := file.Write(data)

	if err != nil {
		log.Fatalf("Error writing init value to file:\n%e", err)
	}
}

func initDefaultData() []byte {
	value := core.CurrencyRates{}

	data, err := json.MarshalIndent(value, "", "\t")

	if err != nil {
		log.Fatalf("Error on marshalling default value\n%e", err)
	}

	return data
}

func (csf *RatesFile) readRatesFile() core.CurrencyRates {
	data := csf.readDataFromBeggining()

	rates := unmarshallRates(data)

	csf.Stored = rates.Stored

	return rates
}

func unmarshallRates(data []byte) core.CurrencyRates {
	rates := core.CurrencyRates{}
	err := json.Unmarshal(data, &rates)

	if err != nil {
		log.Fatalln(err)
	}

	return rates
}

func (rf *RatesFile) clearFile() {
	os.Truncate(rf.path, 0)
}

func (rf *RatesFile) writeRates(rs core.CurrencyRates) {
	data, err := json.MarshalIndent(rs, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	rf.Stored = rs.Stored

	rf.clearFile()
	rf.File.WriteAt(data, 0)
}

func (csf *RatesFile) readDataFromBeggining() []byte {
	csf.File.Seek(0, io.SeekStart)
	data, err := io.ReadAll(csf)

	if err != nil {
		log.Fatalf("Error with reading data from file:\n%e", err)
	}

	return data
}
