package main

import (
	"fmt"
	"strconv"
	"strings"
)

func words(s string) string {
	mots := strings.Fields(s)
	fmt.Println(mots)
	results := []string{}

	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if mot == "(hex)" && len(results) > 0 {
			results[len(results)-1] = hex(results[len(results)-1])
		} else if mot == "(bin)" && len(results) > 0 {
			results[len(results)-1] = bin(results[len(results)-1])
		} else if mot == "(up)" && len(results) > 0 {
			results[len(results)-1] = up(results[len(results)-1])
		} else if mot == "(up," && i+1 < len(mots) {
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombreTexte)
			if err == nil {
				for j := 0; j < n && j < len(results); j++ {
					results[len(results)-1-j] = up(results[len(results)-1-j])
				}
				i++
			}
		} else if mot == "(low)" && len(results) > 0 {
			results[len(results)-1] = low(results[len(results)-1])
		} else if mot == "(low," && i+1 < len(mots) {
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombreTexte)
			if err == nil {
				for j := 0; j < n && j < len(results); j++ {
					results[len(results)-1-j] = low(results[len(results)-1-j])
				}
				i++
			}
		} else if mot == "(cap)" && len(results) > 0 {
			results[len(results)-1] = capitalize(results[len(results)-1])
		} else if mot == "(cap," && i+1 < len(mots) {
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombreTexte)
			if err == nil {
				for j := 0; j < n && j < len(results); j++ {
					results[len(results)-1-j] = capitalize(results[len(results)-1-j])
				}
				i++
			}

		} else {

			results = append(results, mot)

		}

	}
	results = punctuation(results)
	results = article(results)
	results = quote(results)
	result := strings.Join(results, " ")
	return result
}
