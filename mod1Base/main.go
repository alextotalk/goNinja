package main

import (
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
}

func t(b []byte) {
	b = b[1:4]
	b[0] = 64
	fmt.Println(string(b))
}
