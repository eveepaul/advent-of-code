package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// A,X for Rock [1]
// B,Y for Paper [2]
// C,Z for Scissors[3]
// Win [x|y|z] + 6
// Lose 0
// Draw 3

func main() {
	var totalScore int
	combinations := map[string]map[string]int{
		"A": {
			"X": 4,
			"Y": 8,
			"Z": 3,
		},
		"B": {
			"X": 1,
			"Y": 5,
			"Z": 9,
		},
		"C": {
			"X": 7,
			"Y": 2,
			"Z": 6,
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
