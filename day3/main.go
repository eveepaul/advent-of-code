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

    var dup string
    for key := range charCountMap {
      if charCountMap[key] > 0 {
        dup = key
      }
    }
    uppercase := unicode.IsUpper([]rune(dup)[0])
    var charCode int
    caseText := "Lowercase"
    if uppercase {
      caseText = "Uppercase"
      charCode = int(dup[0] - 'A' + 27)
    } else {
      charCode = int(dup[0] - 'a' + 1)
    }
    fmt.Printf("Dup ^ is: %v (%v), code: %v \n", dup, caseText, charCode)
    totalRuck += charCode
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file")
	}
  fmt.Printf("Total ruck: %v \n", totalRuck)
}
