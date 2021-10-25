package main

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("It called before main")
}
func main() {

	fmt.Println("----------------------------")
	massage := "I'll be a ninja soon"
	changeMassage(&massage)
	fmt.Println(&massage)
	fmt.Println("----------------------------")
	number := 5
	var p *int
	p = &number
	fmt.Println("area memory - ", p, "-", "value - ", *p)
	fmt.Println("area memory - ", &number, "-", "value - ", number)
	*p = 10
	fmt.Println("area memory - ", &number, "-", "value - ", number)
	fmt.Println("----------------------------")

	massages := [3]int{}
	fmt.Println(massages)
	massages[0] = 1
	massages[1] = 2
	massages[2] = 3
	fmt.Println(massages)
	printMassages(massages)

}
func printMassages(massages [3]int) error {
	if len(massages) == 0 {
		return errors.New("empty array")
	}
	fmt.Println(massages)
	return nil

}
func changeMassage(massage *string) {
	*massage += " (from function printMassage())"
	fmt.Println(massage)
}
