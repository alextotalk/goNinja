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

	//massages := [3]int{}
	//massages := []string{"1", "2", "3"}
	massages := make([]string, 5)
	//we get panic
	//massages[6]="6"
	strings := append(massages, "6")

	fmt.Println(strings)
	fmt.Println(massages)
	//massages[0] = 1
	//massages[1] = 2
	////massages[2] = 3
	printMassages(massages)

	fmt.Println(massages)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//matrix:= [][]int{}
	matrix := make([][]int, 10)
	count := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			count++
			matrix[x][y] = count
		}
		fmt.Println(matrix[x])
	}
	fmt.Println(count)

	names := []string{
		"Max", "Alex", "Ola",
	}
	for i, v := range names {
		fmt.Println(names[i], "=", v)
	}
	count = 0
	for {
		count++
		if count > 10 {
			break
		}
		fmt.Println(count)
	}
}
func printMassages(massages []string) error {
	if len(massages) == 0 {
		return errors.New("empty array")
	}
	massages[1] = "5"
	fmt.Println(massages)
	return nil

}
func changeMassage(massage *string) {
	*massage += " (from function printMassage())"
	fmt.Println(massage)
}
