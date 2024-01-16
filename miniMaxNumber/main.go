package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 5)
	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4])
	sort.Ints(arr)

	minNumber := arr[0] + arr[1] + arr[2] + arr[3]
	maxNumber := arr[1] + arr[2] + arr[3] + arr[4]

	fmt.Println(minNumber, maxNumber)
}
