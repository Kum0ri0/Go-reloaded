package main

import "testing"

func TestWords(t *testing.T) {
	cas := []struct {
		entree  string
		attendu string
	}{
		{"1E (hex) files were added", "30 files were added"},
	}

	for _, c := range cas {
		obtenu := words(c.entree)
		if obtenu != c.attendu {
			t.Errorf("entrée %q : obtenu %q, attendu %q", c.entree, obtenu, c.attendu)
		}
	}
}