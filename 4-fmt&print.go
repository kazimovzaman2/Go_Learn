package main

import "fmt"

func main() {
	var out string = fmt.Sprintf("Yea: %07d is cool \n", 45) // => store all of it in one
	fmt.Println(out)


	fmt.Printf("Hello %T %v \n", 10, 10) // type, value

	fmt.Printf("Hello %t \n", true) // %t => boolean


	fmt.Printf("Number: %b", 1025) // %b => to binary
	fmt.Printf("Number: %o", 1025) // %o => to octal
	fmt.Printf("Number: %x \n", 1025) // %x => to hexadecimal


	fmt.Printf("NUmber: %e\n", 2.516464444646444) // => scientific representation
	fmt.Printf("Number: %f\n", 2.51616515651651) // => float representation
	fmt.Printf("Number: %g\n", 2.51616516555551) // => represent full



	fmt.Printf("Number: %s\n", "tim") // => represent string
	fmt.Printf("Number: %q\n", "tim") // => represent quotes also


	fmt.Printf("Number: %.2f is cool \n", 3.6464)
	fmt.Printf("Number: %.f is cool \n", 3.6464)
	fmt.Printf("Number: %5.f is cool \n", 3.6464)

	fmt.Printf("Number: %07d is cool \n", 45) // => add 0s

}
