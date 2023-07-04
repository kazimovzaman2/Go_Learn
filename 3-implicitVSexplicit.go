package main

import "fmt"

// implicit - go ozu teyin edir variable i
// explicit - biz edirik

func main()  {
	// var number = 2600.98
	// number = number +5

	number := "hello" // gives automatically var and type
	// var number string = "hello"

	fmt.Printf("%T", number) // print format

	
	var num uint64
	var bl bool
	bl = true

	fmt.Println(num)
	fmt.Println(bl)
}
