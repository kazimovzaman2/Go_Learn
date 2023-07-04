package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Create new variable that scan

	fmt.Printf("Type the year you were born: ")

	scanner.Scan() // Scan start
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // Get input from user, then change to integer and declare to variable named input

	fmt.Printf("You will be %d years old at the end of 2020", 2020 - input)
}
