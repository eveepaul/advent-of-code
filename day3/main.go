package main

import "fmt"

func main() {
	str := "b"
	value := int(str[0] - 'a' + 1)
	fmt.Println(value) // prints: 2

	str = "z"
	value = int(str[0] - 'a' + 1)
	fmt.Println(value) // prints: 26

	str = "A"
	value = int(str[0] - 'A' + 27)
	fmt.Println(value) // prints: 27

	str = "Z"
	value = int(str[0] - 'A' + 27)
	fmt.Println(value) // prints: 52
}
