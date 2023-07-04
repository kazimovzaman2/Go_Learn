package main

import (
	"fmt"
)


type Student struct {
	name string
	grades []int
	age int
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s *Student) getAvarageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	currMax := 0

	for _, v := range s.grades {
		if currMax < v {
			currMax = v
		}
	}

	return currMax
}



func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	// average := s1.getAvarageGrade()
	// fmt.Println(average)


	maxGrade := s1.getMaxGrade()
	fmt.Println(maxGrade)

}
