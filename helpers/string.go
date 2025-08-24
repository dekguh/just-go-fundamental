package helpers

import "strings"

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
