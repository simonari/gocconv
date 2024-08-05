package core

import (
	"log"
	"slices"
	"testing"
)

func TestNewRateFunc(t *testing.T) {
	c := NewRate("a", "b", 1)

	expected := struct {
		From string
		To   string
		Rate float32
	}{
		From: "A",
		To:   "B",
		Rate: 1,
	}

	if c.From != expected.From {
		t.Fatalf("Expected %s\nGot: %s\n", expected.From, c.From)
	}
	if c.To != expected.To {
		t.Fatalf("Expected %s\nGot: %s\n", expected.To, c.To)
	}
	if c.Rate != expected.Rate {
		t.Fatalf("Expected %.2f\nGot: %.2f\n", expected.Rate, c.Rate)
	}
}

func TestReverseRateFromIntegerValue(t *testing.T) {
	c := NewRate("A", "B", 2)

	rc := c.ReverseRate()

	if c.From != rc.To {
		t.Fatalf("Expected %s\nGot: %s\n", c.From, rc.To)
	}
	if c.To != rc.From {
		t.Fatalf("Expected %s\nGot: %s\n", c.To, rc.From)
	}
	if 1/c.Rate != rc.Rate {
		t.Fatalf("Expected %.2f\nGot: %.2f\n", 1/c.Rate, rc.Rate)
	}
}

func TestReverseRateFromFloatValue(t *testing.T) {
	c := NewRate("A", "B", 0.33)

	rc := c.ReverseRate()

	if c.From != rc.To {
		t.Fatalf("Expected %s\nGot: %s\n", c.From, rc.To)
	}
	if c.To != rc.From {
		t.Fatalf("Expected %s\nGot: %s\n", c.To, rc.From)
	}
	if 1/c.Rate != rc.Rate {
		t.Fatalf("Expected %.2f\nGot: %.2f\n", 1/c.Rate, rc.Rate)
	}
}

func TestNewCurrencyRatesFunc(t *testing.T) {
	cs := make([]CurrencyRate, 3)

	cs[0] = *NewRate("a", "b", 1)
	cs[1] = *NewRate("b", "c", 1)
	cs[2] = *NewRate("c", "d", 1)

	crs := NewCurrencyRates(cs, uint8(len(cs)))

	expected := struct {
		Rates  []CurrencyRate
		Stored uint8
	}{
		Rates: []CurrencyRate{
			{"A", "B", 1},
			{"B", "C", 1},
			{"C", "D", 1},
		},
		Stored: 3,
	}

	if !slices.Equal(expected.Rates, crs.Rates) {
		log.Fatalf("'Rates' field formed bad")
	}

	if expected.Stored != crs.Stored {
		log.Fatalf("'Stored' fields formed bad")
	}
}

func TestForwardGetRateIdxMethod(t *testing.T) {
	cs := make([]CurrencyRate, 3)
	cs[1] = *NewRate("A", "B", 1)

	crs := NewCurrencyRates(cs, 1)

	idx, err := crs.getRateIdx(cs[1])

	if err != nil {
		t.Fatalf("%s", err)
	}

	if idx != 1 {
		t.Fatalf("Wrong index:\nGot: %v\nExpected: %v", idx, 1)
	}
}

func TestReverseGetRateIdxMethod(t *testing.T) {
	cs := make([]CurrencyRate, 3)
	cs[1] = *NewRate("A", "B", 1)

	crs := NewCurrencyRates(cs, 1)

	idx, err := crs.getRateIdx(cs[1].ReverseRate())

	if err != nil {
		t.Fatalf("%s", err)
	}

	if idx != 1 {
		t.Fatalf("Wrong index:\nGot: %v\nExpected: %v", idx, 1)
	}
}

func TestAddRateMethod(t *testing.T) {
	rates := NewCurrencyRates([]CurrencyRate{}, 0)

	rate := NewRate("a", "b", 1)
	rates.Add(rate)

	if rates.Stored != 1 {
		t.Fatalf("Wrong Stored value:\n Got: %v\nExcpected: %v", rates.Stored, 1)
	}

	if len(rates.Rates) != 1 {
		t.Fatalf("Wrong length:\nGot: %v\nExpected: %v", len(rates.Rates), 1)
	}
}

func TestGetForwardRateMethod(t *testing.T) {
	rates := NewCurrencyRates(
		[]CurrencyRate{
			*NewRate("a", "b", 1),
			*NewRate("b", "c", 1),
		},
		2,
	)

	expected := struct {
		From string
		To   string
		Rate float32
	}{
		From: "A", To: "B", Rate: 1.0,
	}

	rate := rates.Get("A", "b")

	if expected.From != rate.From {
		t.Fatalf("Wrong From field:\nExpected: %v\nGot: %v\n", expected.From, rate.From)
	}
	if expected.To != rate.To {
		t.Fatalf("Wrong To field:\nExpected: %v\nGot: %v\n", expected.To, rate.To)
	}
	if expected.Rate != rate.Rate {
		t.Fatalf("Wrong Rate field:\nExpected: %v\nGot: %v\n", expected.Rate, rate.Rate)
	}
}

func TestGetReverseRateMethod(t *testing.T) {
	rates := NewCurrencyRates(
		[]CurrencyRate{
			*NewRate("a", "b", 1),
			*NewRate("b", "c", 1),
		},
		2,
	)

	expected := struct {
		From string
		To   string
		Rate float32
	}{
		From: "B", To: "A", Rate: 1.0,
	}

	rate := rates.Get("b", "a")

	if expected.From != rate.From {
		t.Fatalf("Wrong From field:\nExpected: %v\nGot: %v\n", expected.From, rate.From)
	}
	if expected.To != rate.To {
		t.Fatalf("Wrong To field:\nExpected: %v\nGot: %v\n", expected.To, rate.To)
	}
	if expected.Rate != rate.Rate {
		t.Fatalf("Wrong Rate field:\nExpected: %v\nGot: %v\n", expected.Rate, rate.Rate)
	}
}

func TestGetRateMethodReturnsNil(t *testing.T) {
	rates := NewCurrencyRates([]CurrencyRate{}, 0)

	rate := rates.Get("A", "B")

	if rate != nil {
		t.Fatalf("Wrong value:\nExpected: %v\nGot: %v\n", nil, rate)
	}
}

func TestUpdateRateMethod(t *testing.T) {
	rates := NewCurrencyRates(
		[]CurrencyRate{
			*NewRate("A", "B", 1),
		},
		1,
	)

	expected := struct {
		From string
		To   string
		Rate float32
	}{
		"A", "B", 2,
	}

	err := rates.Update(expected.From, expected.To, expected.Rate)

	if err != nil {
		t.Fatalf("Error occured during update:\n%s\n", err)
	}

	rate := rates.Get(expected.From, expected.To)

	if rate == nil {
		t.Fatalf("Rate dissapeared")
	}

	if expected.Rate != rate.Rate {
		t.Fatalf("Wrong Rate value:\nExpected: %v\nGot: %v\n", expected.Rate, rate.Rate)
	}
}

func TestDeleteRateMethod(t *testing.T) {
	rates := NewCurrencyRates(
		[]CurrencyRate{
			*NewRate("A", "B", 1),
			*NewRate("B", "C", 2),
			*NewRate("C", "D", 3),
		},
		1,
	)

	expected := struct {
		Rates  []CurrencyRate
		Stored uint8
	}{
		[]CurrencyRate{
			*NewRate("B", "C", 2),
		},
		1,
	}

	toDelete := []struct {
		From string
		To   string
	}{
		{"A", "B"},
		{"C", "D"},
	}

	for _, item := range toDelete {
		rates.Delete(item.From, item.To)
	}

	if expected.Stored != rates.Stored {
		t.Fatalf("Wrong Stored value:\nExpected: %v\nGot: %v\n", expected.Stored, rates.Stored)
	}

	if !slices.Equal(rates.Rates, expected.Rates) {
		t.Fatalf("Wrong Stored value:\nExpected: %v\nGot: %v\n", expected.Rates, rates.Rates)
	}
}
