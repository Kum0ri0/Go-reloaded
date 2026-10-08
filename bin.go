package main

import (
	"fmt"
	"strconv"
)

func bin(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("Erreur dans la Convertion bin :", err)
		return s
	}
	return strconv.FormatInt(n, 10)

}
