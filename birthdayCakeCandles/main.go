package main

import (
	"fmt"
)

func main() {
	var length, temp, max, occ int
	fmt.Scan(&length)

	a := []int{}
	for i := 0; i < length; i++ {
		fmt.Scan(&temp)
		a = append(a, temp)
		if temp > max {
			max = temp
		}
	}

	for _, num := range a {
		if num == max {
			occ++
		}
	}
	fmt.Println(occ)
}
