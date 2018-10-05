package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func worldCount(s string) (result map[string]int) {
	fields := strings.Fields(s)
	result = make(map[string]int, len(fields))
	for _, value := range fields {
		result[value]++
	}
	return
}

func main() {
	wc.Test(worldCount)
}
