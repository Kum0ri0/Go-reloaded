package main

func quote(s []string) []string {
	results := []string{}
	ouvert := false
	aColler := false

	for i := 0; i < len(s); i++ {
		if s[i] == "'" && !ouvert {
			ouvert = true
			aColler = true
		} else if s[i] == "'" && ouvert && len(results) > 0 {
			results[len(results)-1] = results[len(results)-1] + "'"
			ouvert = false
		} else {
			mot := s[i]
			if aColler == true {
				mot = "'" + mot
				aColler = false
			}
			results = append(results, mot)
		}
	}
	return results
}
