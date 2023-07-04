package main

import (
	"fmt"
)

type Point struct {
	x int32
	y int32
}


func changeX(pt *Point) {
	pt.x = 100
}


func main() {

	// var p1 Point = Point{1, 2}
	// p1 := Point{1, 2}

	// p1.x = 7

	p1 := &Point{y: 2}

	fmt.Println(p1)
	changeX(p1)
	fmt.Println(p1)	

}
