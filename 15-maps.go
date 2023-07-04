package main

import (
	"fmt"
)

func main() {

	var mp map[string]int = map[string]int {
		"apple": 5,
		"pear": 6,
		"orange": 9,
	}


	// mp["tim"] = 900   // create 
	// delete(mp, "apple")   // delete


	// val, ok := mp["apple"]   // val = key, ok = is there like this staf
	// fmt.Println(val,ok)


	fmt.Println(len(mp))
	

}
