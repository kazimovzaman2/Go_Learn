package main

import (
	"fmt"
)

func main() {

	var a[]int = []int{1,5,2,3,4,5,4,6,323}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if (element == element2) {
				fmt.Println(element)
			}
		}
	}

}
