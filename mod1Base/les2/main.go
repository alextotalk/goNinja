package main

import (
	"errors"
	"fmt"
	"log"
)

var m, k int = 7, 8

func main() {

	massager1 := sayHello("Alex", 30)
	printMassager(massager1)
	//printMassager("text 1")
	//printMassager("text 2")
	//printMassager("text 3")
	massager, err := anterTheClub(17)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(massager)
	//fmt.Println(anter)
	//fmt.Println(anterTheClub(65))
	//fmt.Println(anterTheClub(64))
	//fmt.Println(anterTheClub(17))
}
func printMassager(massager string) {

	fmt.Println(massager)
}
func sayHello(name string, age int) string {

	return fmt.Sprintf("Hi, %s !!! You %d years old. :)", name, age)
}

func anterTheClub(age int) (string, error) {
	if age >= 18 && age < 65 {
		return "Enter, but careful", nil
	} else if age >= 65 {
		return "Stop, You are very old", errors.New("You are very old")
	}

	//else {
	//return "Stop, You are very young", false
	//}
	return "Stop, You are very young", errors.New("You are very young")
}
