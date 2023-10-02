package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// str := "b"
	// value := int(str[0] - 'a' + 1)
	// fmt.Println(value) // prints: 2

	// str = "z"
	// value = int(str[0] - 'a' + 1)
	// fmt.Println(value) // prints: 26

	// str = "A"
	// value = int(str[0] - 'A' + 27)
	// fmt.Println(value) // prints: 27

	// str = "Z"
	// value = int(str[0] - 'A' + 27)
	// fmt.Println(value) // prints: 52

	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal("Error reading input file")
	}

	scanner := bufio.NewScanner(file)
	// var totalRuck int

	for scanner.Scan() {

		str := scanner.Text()
		lastHalf := str[len(str)/2:]
		firstHalf := str[:len(str)/2]
		fmt.Printf("Line: {%v} {%v} {%v} \n", str, firstHalf, lastHalf)

		charCountMap := make(map[string]int)
		for i := 0; i < len(firstHalf); i++ {
			_, exists := charCountMap[string(firstHalf[i])]
			if exists {
				charCountMap[string(firstHalf[i])] += 1
			} else {
				charCountMap[string(firstHalf[i])] = 0
			}
		}

		for key := range charCountMap {
			charCountMap[key] = 0
		}

		for i := 0; i < len(lastHalf); i++ {
			_, exist := charCountMap[string(lastHalf[i])]
			if exist {
				charCountMap[string(lastHalf[i])] += 1
			}
		}

		// TODO: get the key of non zero value in the map
		// get the unicode value
		// add to totalRuck
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file")
	}
}
