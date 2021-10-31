package main

import (
	"errors"
	"fmt"
)

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(string(b))

	// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
	t(b)
	fmt.Println(string(b))

	s := []byte{65, 45, 122} // a slice referencing the storage of x
	fmt.Println("s=", string(s))
	t1 := make([]byte, len(s)+1, (cap(s)+1)*2)
	fmt.Println("t1=", string(t1), "len=", len(t1), "cap=", cap(t1))

	copy(t1, s)
	s = t1

	t1 = append(t1, 64)
	fmt.Println("s=", string(s))
	fmt.Println("t1=", string(t1), "len=", len(t1), "cap=", cap(t1))
	st1 := append(s, t1...)
	fmt.Println(st1)
	massages := []string{"1", "2", "3"}
	fmt.Println(massages)
	//massages[0] = 1
	//massages[1] = 2
	////massages[2] = 3
	printMassages(&massages)

	fmt.Println(massages)
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
	slice := make([]byte, 50)

	fmt.Println("Before: len(slice) =", len(slice))
	//newSlice :=
	SubtractOneFromLength(&slice)
	fmt.Println("After:  len(slice) =", len(slice), slice)
	//fmt.Println("After:  len(newSlice) =", len(newSlice))
	// Make a copy of a slice (of int).
	slices := []int{32, 43, 23}
	slice3 := append([]int(nil), slices...)
	fmt.Println("Copy a slice:", slice3)
}
func SubtractOneFromLength(slice *[]byte) {
	pSlice := *slice
	*slice = pSlice[0 : len(pSlice)-1]
}
func printMassages(massages *[]string) error {
	if len(*massages) == 0 {
		return errors.New("empty array")
	}
	msg := *massages
	msg[1] = "5"
	msg = append(msg, "6")
	massages = &msg
	fmt.Println(massages)
	return nil

}

func t(b []byte) {
	b = b[1:4]
	b[0] = 64
	fmt.Println(string(b))
}
func Min(num ...int) (int, error) {
	if len(num) == 0 {
		return 0, errors.New("not numbers")
	}
	min := num[0]

	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min, nil
}
