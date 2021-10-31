package main

import (
	"fmt"
)

func main() {
	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 48,
	}
	delete(users, "Vasya")
	age, exists := users["asd"]
	if !exists {
		fmt.Println("not exist")
	}
	fmt.Println(age)
	for key, value := range users {
		fmt.Println(key, value)
	}
	//var panicInMap map[string]int
	//initialization map
	someMap := make(map[string]int)
	////we get panic
	//panicInMap["Vasya"]=2
	someMap["1"] = 10
	someMap["2"] = 20
	fmt.Println("length of map someMap =", len(someMap))
}
