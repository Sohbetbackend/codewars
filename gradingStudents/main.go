package main

import "fmt"

func main() {
	var n, grade int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&grade)
		if grade < 38 {
			fmt.Println(grade)
		} else {
			if 5-(grade%5) < 3 {
				fmt.Println(grade + 5 - (grade % 5))
			} else {
				fmt.Println(grade)
			}
		}
	}
}
