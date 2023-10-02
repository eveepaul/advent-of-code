package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// X means you need to lose,
// Y means you need to end the round in a draw,
// and Z means you need to win

func main() {
	var totalScore int
	combinations := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 4,
			"Z": 8,
		},
		"B": {
			"X": 1,
			"Y": 5,
			"Z": 9,
		},
		"C": {
			"X": 2,
			"Y": 6,
			"Z": 7,
		},
	}

	labelCombi := map[string]string{
		"A": "Rock",
		"X": "Rock",
		"B": "Paper",
		"Y": "Paper",
		"C": "Scissors",
		"Z": "Scissors",
	}

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal("Error in reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())
		score := combinations[fields[0]][fields[1]]
		totalScore = totalScore + score
		fmt.Printf("%v vs %v -> Score: %v, total: %v \n", labelCombi[fields[0]], labelCombi[fields[1]], score, totalScore)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanner")
	}
}
