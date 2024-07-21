package core

import (
	"errors"
	"fmt"
	"log"
)

type CurrencyRate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float32 `json:"rate"`
}

func (cer *CurrencyRate) ReverseRate() CurrencyRate {
	return CurrencyRate{
		From: cer.To,
		To:   cer.From,
		Rate: 1 / cer.Rate,
	}
}

type CurrencyRates struct {
	Rates  []CurrencyRate `json:"rates"`
	Stored uint8          `json:"stored"`
}

func (rs *CurrencyRates) getRateIdx(item CurrencyRate) (int, error) {
	for i, r := range rs.Rates {
		forward := (r.From == item.From && r.To == item.To)
		reverse := (r.From == item.To && r.To == item.From)

		if forward || reverse {
			return i, nil
		}
	}

	return -1, errors.New("no rate found")
}

func (rs *CurrencyRates) Add(r CurrencyRate) []CurrencyRate {
	_, err := rs.getRateIdx(r)

	if err != nil {
		log.Fatalf("[!] Error: %s", err)
	}

	result := rs.Rates
	result = append(result, r)

	return result
}

func (rs *CurrencyRates) Get(from, to string) *CurrencyRate {
	idx, err := rs.getRateIdx(CurrencyRate{From: from, To: to})

	if err != nil {
		fmt.Printf("[!] %s", err)
		return nil
	}

	return &rs.Rates[idx]
}

func (rs *CurrencyRates) Update(from, to string, rate float32) error {
	r := CurrencyRate{From: from, To: to, Rate: rate}

	idx, err := rs.getRateIdx(r)

	if err != nil {
		return err
	}

	rs.Rates[idx] = r

	return nil
}

func (rs *CurrencyRates) Delete(from, to string) error {
	rateToDelete := CurrencyRate{From: from, To: to}
	idx, err := rs.getRateIdx(rateToDelete)

	if err != nil {
		return err
	}

	result := make([]CurrencyRate, len(rs.Rates)-1)

	for i := 0; i < idx; i++ {
		result[i] = rs.Rates[i]
	}

	for i := idx + 1; i < len(rs.Rates); i++ {
		result[i] = rs.Rates[i]
	}

	rs.Rates = result
	rs.Stored = uint8(len(rs.Rates) - 1)

	return nil
}
