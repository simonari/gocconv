package storage

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"vsimonari/gocconv/internal/core"
)

func (rf *RatesFile) ReadAll() core.CurrencyRates {
	rf.File.Seek(0, io.SeekStart)
	data, err := io.ReadAll(rf)

	if err != nil {
		log.Fatalf("[!] Error: %s\n", err)
	}

	rates := core.CurrencyRates{}

	if err := json.Unmarshal(data, &rates); err != nil {
		log.Fatalln(err)
	}

	rf.Stored = rates.Stored

	return rates
}

func (rf *RatesFile) Write(rs core.CurrencyRates) {
	defer rf.Close()

	data, err := json.MarshalIndent(rs, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	rf.Stored = rs.Stored

	os.Truncate(rf.path, 0)
	rf.File.WriteAt(data, 0)
}
