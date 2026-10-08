package main

import (
	"strings"
)

func article(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if strings.Contains("aeiouhAEIOUH", s[i+1][:1]) {
				if s[i] == "a" {
					s[i] = "an"
				}
				if s[i] == "A" {
					s[i] = "An"
				}
			}
		}
	}
	return s
}
