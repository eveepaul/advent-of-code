package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
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
			char := getMaxChar(elvesGroup)
			charCode := getCode(char)
			totalRuck = totalRuck + charCode
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file")
	}
	fmt.Printf("Total ruck: %v \n", totalRuck)
}

func getMaxChar(strGroup []string) rune {
	var charMaps []map[rune]int
	for _, str := range strGroup {
		charMap := make(map[rune]int)
		for _, s := range str {
			if _, exist := charMap[s]; exist {
				charMap[s]++
			} else {
				charMap[s] = 1
			}
		}
		charMaps = append(charMaps, charMap)
	}

	maxChar := make(map[rune]int)
	for key, val := range charMaps[0] {

		key2, exist2 := charMaps[1][key]
		key3, exist3 := charMaps[2][key]

		if exist2 && exist3 {
			maxChar[key] = val + key2 + key3
		}
	}
	var r rune
	for key, _ := range maxChar {
		r = key
		break
	}
	return r
}

func getCode(r rune) int {
	uppercase := unicode.IsUpper(r)
	var charCode int
	if uppercase {
		charCode = int(r - 'A' + 27)
	} else {
		charCode = int(r - 'a' + 1)
	}
	return charCode
}
