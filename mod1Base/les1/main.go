package main

import (
	"fmt"
	"reflect"
)

var m, k int = 7, 8

func main() {
	//const massage string = "Я скоро буду Гофером"
	var massage string = "Я скоро буду Гофером"
	massage = "I am a ninja"
	var nameVar = "some text"
	var a bool
	var b rune = 'a'

	fmt.Println(massage, nameVar)
	fmt.Println(reflect.TypeOf(massage))
	fmt.Println("/n", a)
	fmt.Println("/n", b)

	m, k = 2, 3
	fmt.Println(m, k)

	m, k = k, m
	fmt.Println(m, k)
	print()
}
func print() {
	//var m, k int = 4, 5
	fmt.Println(m, k)
}
