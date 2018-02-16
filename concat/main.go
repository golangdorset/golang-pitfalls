package main

import "fmt"

func main() {
	fmt.Println("H" + "i")
	// Output: Hi

	fmt.Println('H' + 'i') // HL
	// Output: 177

	fmt.Printf("%T", 'i')
}
