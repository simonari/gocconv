package core

import (
	"errors"
	"fmt"
	"strings"
)

type CurrencyRate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float32 `json:"rate"`
}

func NewRate(from, to string, rate float32) *CurrencyRate {
	return &CurrencyRate{
		From: strings.ToUpper(from),
		To:   strings.ToUpper(to),
		Rate: rate,
	}
}

func (r *CurrencyRate) ReverseRate() CurrencyRate {
	return CurrencyRate{
		From: r.To,
		To:   r.From,
		Rate: 1 / r.Rate,
	}
}

type CurrencyRates struct {
	Rates  []CurrencyRate `json:"rates"`
	Stored uint8          `json:"stored"`
}

func NewCurrencyRates(rates []CurrencyRate, stored uint8) *CurrencyRates {
	return &CurrencyRates{
		Rates:  rates,
		Stored: stored,
	}
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

func (rs *CurrencyRates) Add(r *CurrencyRate) []CurrencyRate {
	_, err := rs.getRateIdx(*r)

	result := rs.Rates

	if err != nil {
		result = append(result, *r)
	}

	rs.Rates = result
	rs.Stored = uint8(len(result))

	return result
}

func (rs *CurrencyRates) Get(from, to string) *CurrencyRate {
	from, to = capitilizeFromTo(from, to)

	idx, err := rs.getRateIdx(CurrencyRate{From: from, To: to})

	if err != nil {
		fmt.Printf("[!] %s\n", err)
		return nil
	}

	r := rs.Rates[idx]

	if from == r.From {
		return &r
	} else {
		r = r.ReverseRate()
		return &r
	}
}

func (rs *CurrencyRates) Update(from, to string, rate float32) error {
	from, to = capitilizeFromTo(from, to)

	r := CurrencyRate{From: from, To: to, Rate: rate}

	idx, err := rs.getRateIdx(r)

	if err != nil {
		fmt.Printf("[!] %s\n", err)
		return err
	}

	rs.Rates[idx] = r

	return nil
}

func (rs *CurrencyRates) Delete(from, to string) error {
	from, to = capitilizeFromTo(from, to)

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
		result[i-1] = rs.Rates[i]
	}

	rs.Stored = uint8(len(rs.Rates) - 1)
	rs.Rates = result

	return nil
}

func capitilizeFromTo(from, to string) (string, string) {
	return strings.ToUpper(from), strings.ToUpper(to)
}
