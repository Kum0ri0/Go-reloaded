package main

import (
	"fmt"
	"strings"
)

func words(s string) string {
	mots := strings.Fields(s)
	fmt.Println(mots)
	results := []string{}

	for _, mot := range mots {
		if mot == "(hex)" && len(results) > 0 {
			results[len(results)-1] = hex(results[len(results)-1])
		} else if mot == "(bin)" && len(results) > 0 {
			results[len(results)-1] = bin(results[len(results)-1])
		} else if mot == "(up)" && len(results) > 0 {
			results[len(results)-1] = up(results[len(results)-1])
		} else if mot == "(low)" && len(results) > 0 {
			results[len(results)-1] = low(results[len(results)-1])
		} else if mot == "(cap)" && len(results) > 0 {
			results[len(results)-1] = capitalize(results[len(results)-1])
		} else {
			results = append(results, mot)

		}

	}

	result := strings.Join(results, " ")
	return result
}
