package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println(Min(2, 3, 443, -32, -1))
	fmt.Println(Max(2, 3, 443, -32, -1))
	// Min()
	// massager, err := enterTheClub(18)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(massager)
	res, err := enterTheClub(22)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(prediction("day"))
	func() {
		fmt.Println("anonymous funk")
	}()
	anonymousFunk := func() {
		fmt.Println("anonymous funk")
	}
	fmt.Println("It is also")
	anonymousFunk()
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func increment2() int {
	count := 0
	count++
	return count
}
func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 65 {
		return "Enter, but careful", nil
	} else if age >= 65 {
		return "Stop, You are very old", errors.New("you are very old")
	}

	return "Stop, You are very young", errors.New("you are very young")
}

func prediction(dayOfWeek string) (string, error) {
	//if dayOfWeek == "mn"{
	//	return "Have a good " + dayOfWeek
	//} else if dayOfWeek  =="tu" {
	//	return "Have a good " + dayOfWeek
	//} else if dayOfWeek  =="wn" {
	//	return "Have a good " + dayOfWeek
	//} else if dayOfWeek  =="th" {
	//	return "Have a good " + dayOfWeek
	//}
	switch dayOfWeek {
	case "mn":
		return "Have a good " + dayOfWeek, nil
	case "tu":
		return "Have a good " + dayOfWeek, nil
	case "wn":
		return "Have a good " + dayOfWeek, nil
	case "th":
		return "Have a good " + dayOfWeek, nil
	default:
		return "invalid write", errors.New("invalid day of the week")
	}
}

func Min(num ...int) int {
	if len(num) == 0 {
		return 0
	}
	min := num[0]
	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}
func Max(num ...int) int {
	if len(num) == 0 {
		return 0
	}
	max := num[0]
	for _, v := range num {
		if v > max {
			max = v
		}
	}
	return max
}
