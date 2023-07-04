package main

import (
	"fmt"
)

func main() {

	// var arr [5]int

	// arr[0] = 100
	// arr[4] = 900

	// fmt.Println(arr[0])

	arr := [3]int{4,5,6}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

}
