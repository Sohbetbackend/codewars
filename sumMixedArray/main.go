package main

import (
	"strconv"
)

func SumMix(arr []any) int {
	sum := 0
	for _, num := range arr {
		switch num := num.(type) {
		case int:
			sum += num
		case string:
			k, _ := strconv.Atoi(num)
			sum += k
		}
	}
	return sum
}

func main() {
	SumMix([]any{9, 3, "7", "3"})
}
