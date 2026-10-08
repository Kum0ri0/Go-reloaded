package main

import (
	"fmt"
	"strconv"
)

func hex(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("Erreur dans la Convertion hex :", err)
		return s
	}
	return strconv.FormatInt(n, 10)

}
