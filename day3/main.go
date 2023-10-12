package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "unicode"
)

func main() {

	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal("Error reading input file")
	}

	scanner := bufio.NewScanner(file)
	var totalRuck int
	elvesGroup := make([]string, 3)
	var elvesCount int
	for scanner.Scan() {

		str := scanner.Text()
		elvesGroup[elvesCount] = str
		elvesCount++
		if elvesCount == 3 {

			elvesCount = 0

			fmt.Printf("Elves Group: %v \n", elvesGroup)

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file")
	}
	fmt.Printf("Total ruck: %v \n", totalRuck)
}
