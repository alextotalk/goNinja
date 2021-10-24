package main

import (
	"errors"
	"fmt"
 )

func main() {
	fmt.Println(Min(2))
	// Min()
	// massager, err := enterTheClub(18)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(massager)

	fmt.Println(prediction("dwedw"))
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 65 {
		return "Enter, but careful", nil
	} else if age >= 65 {
		return "Stop, You are very old", errors.New("You are very old")
	}

	return "Stop, You are very young", errors.New("You are very young")
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
