package core

import (
	"errors"
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

type ratePresented struct {
	forward bool
	reverse bool
}

func (rp *ratePresented) isBoth() bool {
	if rp.forward && rp.reverse {
		return true
	}

	return false
}

type CurrencyRates struct {
	Rates  []CurrencyRate `json:"rates"`
	Stored uint8          `json:"stored"`
}

func (rs *CurrencyRates) isRatePresented(item CurrencyRate) ratePresented {
	rp := ratePresented{}

	for _, cer := range rs.Rates {
		if !rp.forward {
			rp.forward = bool(cer.From == item.From && cer.To == item.To)
		}

		if !rp.reverse {
			rp.reverse = bool(cer.From == item.To && cer.To == item.From)
		}
	}

	return rp
}

func (rs *CurrencyRates) Add(item CurrencyRate) []CurrencyRate {
	result := rs.Rates

	isPresented := rs.isRatePresented(item)

	if !isPresented.forward {
		result = append(result, item)
	}

	if !isPresented.reverse {
		result = append(result, item.ReverseRate())
	}

	return result
}

func (rs *CurrencyRates) Get(from, to string) *CurrencyRate {
	for _, r := range rs.Rates {
		if r.From == from && r.To == to {
			return &r
		}
	}

	return nil
}

func (rs *CurrencyRates) Update(from, to string, rate float32) error {
	updatedCount := 0

	for i, r := range rs.Rates {
		if r.From == from && r.To == to {
			rs.Rates[i].Rate = rate
			updatedCount += 1
		}

		if r.From == to && r.To == from {
			rs.Rates[i].Rate = 1 / rate
			updatedCount += 1
		}
	}

	if updatedCount != 0 {
		return nil
	}

	println(updatedCount)

	return errors.New("rate was not found")
}

func (rs *CurrencyRates) Delete(from, to string) error {
	isPresented := ratePresented{}
	var toRemoveForwardIdx, toRemoveReverseIdx int

	for i, r := range rs.Rates {
		if r.From == from && r.To == to {
			isPresented.forward = true
			toRemoveForwardIdx = i
		}

		if r.From == to && r.To == from {
			isPresented.reverse = true
			toRemoveReverseIdx = i
		}
	}

	if isPresented.isBoth() {
		result := make([]CurrencyRate, len(rs.Rates)-2)

		var lesserIdx, greaterIdx int

		if toRemoveForwardIdx < toRemoveReverseIdx {
			lesserIdx, greaterIdx = toRemoveForwardIdx, toRemoveReverseIdx
		} else {
			lesserIdx, greaterIdx = toRemoveReverseIdx, toRemoveForwardIdx
		}

		for i := 0; i < lesserIdx; i++ {
			result[i] = rs.Rates[i]
		}

		for i := lesserIdx + 1; i < greaterIdx; i++ {
			result[i] = rs.Rates[i]
		}

		for i := greaterIdx + 1; i < len(rs.Rates)-2; i++ {
			result[i] = rs.Rates[i]
		}

		rs.Rates = result
		rs.Stored = uint8(rs.Stored - 2)

		return nil

	} else if isPresented.forward {
		return rs.deleteSingle(toRemoveForwardIdx)
	} else if isPresented.reverse {
		return rs.deleteSingle(toRemoveForwardIdx)
	} else {
		return errors.New("can not find the following element")
	}
}

func (rs *CurrencyRates) deleteSingle(idx int) error {
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
