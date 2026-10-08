package main

import (
	"strings"
)

func punctuation(s []string) []string {
	results := []string{}
	for i := 0; i < len(s); i++ {
		reste := strings.TrimLeft(s[i], ".,!?:;")
		ponct := s[i][:len(s[i])-len(reste)]

		if ponct != "" && len(results) > 0 {

			results[len(results)-1] = results[len(results)-1] + ponct
		}
		if reste != "" {
			results = append(results, reste)
		}
	}
	return results
}
