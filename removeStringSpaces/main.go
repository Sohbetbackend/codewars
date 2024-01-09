package main

import (
	"strings"
)

func NoSpace(word string) string {
	word = strings.ReplaceAll(word, " ", "")
	return word
}
