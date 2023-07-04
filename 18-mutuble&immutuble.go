package main

import (
	"fmt"
)

func changeFirst(slice []int) {
	slice[0] = 1000
}


func main() {

	// var x []int = []int{3,4,5}
	// y := x
	// y[0] = 100

	// var x map[string]int = map[string]int{"hello": 3}
	// y := x
	// y["y"] = 100
	// x["7"] = 7

	// var x[2]int = [2]int {3,4}
	// y := x
	// y[0] = 100

	var x[]int = []int{3,4,5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)

	
}
