package main

import "fmt"

func main() {
	a := []byte("ab")

	a1 := append(a, 'c') // HL
	a2 := append(a, 'd') // HL

	fmt.Println(string(a1))
	// Output: abd

	fmt.Println(string(a2))
	// Output: abd
}
