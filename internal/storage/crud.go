package storage

import (
	"fmt"
	"log"
	"vsimonari/gocconv/internal/core"
)

func (rf *RatesFile) AddRate(newRate core.CurrencyRate) {
	rs := rf.readRatesFile()

	rates := rs.Add(newRate)

	rs = core.CurrencyRates{Rates: rates, Stored: uint8(len(rates))}

	fmt.Println("[+] rate added")

	rf.writeRates(rs)
}

func (rf *RatesFile) GetRate(from, to string) *core.CurrencyRate {
	rs := rf.readRatesFile()

	return rs.Get(from, to)
}

func (rf *RatesFile) UpdateRate(from, to string, r float32) {
	rs := rf.readRatesFile()

	err := rs.Update(from, to, r)

	if err != nil {
		log.Fatalf("[!] Error: %s", err)
	}

	rf.writeRates(rs)
}

func (rf *RatesFile) DeleteRate(from, to string) {
	rs := rf.readRatesFile()

	err := rs.Delete(from, to)

	if err != nil {
		log.Fatalf("[!] Error: %s", err)
	}

	rf.writeRates(rs)
}
