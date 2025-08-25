package helpers

import (
	"fmt"
	"strings"
)

func TextCapitalize(words string) string {
	splitWords := strings.Split(words, " ")
	tempData := []string{}

	for _, word := range splitWords {
		firstLetter := strings.ToUpper(word[0:1])
		restWord := word[1:]
		tempData = append(tempData, firstLetter+restWord)
	}

	return strings.Join(tempData, " ")
}

func GenerateCarDetail() string {
	car := Vehicle{
		brand: "Toyota",
		model: "Avanza",
		Engine: Engine{
			cc:   1500,
			fuel: "Bensin",
		},
	}

	return fmt.Sprintf("Brand: %s, Model: %s, Engine: %d cc, Fuel: %s", car.brand, car.model, car.Engine.cc, car.Engine.fuel)
}
