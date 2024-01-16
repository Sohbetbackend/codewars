package main

import "fmt"

func main() {
    var sum int
    var array []int{1,2,3,4,10,11}
    for i := range array{
        sum = sum + array[i]
    }
    fmt.Println(sum)     
}
