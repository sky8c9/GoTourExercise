// Exercise: Maps

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wMap := make(map[string]int)
	words := strings.Fields(s)
	
	for _, word := range words {
		cnt := wMap[word]	
		wMap[word] = cnt + 1
	}
	
	return wMap
}

func main() {
	wc.Test(WordCount)
}