package main

import "fmt"

func main() {
	defer handlerPanic()
	//names := []string{
	//	"Max", "Alex", "Ola", "Vas",
	//}
	//names[4] = "Bob"
	//fmt.Println(names)
	panic("HElP")
}
func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("handlerPanic is done")
}
