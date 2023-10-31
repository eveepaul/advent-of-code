package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	lowBound  int
	highBound int
}

func main() {
	fmt.Println("hello")
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal("Error reading input file")
	}

	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")

		pair1Low, err := strconv.Atoi(first[0])
		if err != nil {
			log.Fatal("Invalid string to convert")
		}
		pair1High, err := strconv.Atoi(first[1])
		if err != nil {
			log.Fatal("Invalid string to convert")
		}
		pair2Low, err := strconv.Atoi(second[0])
		if err != nil {
			log.Fatal("Invalid string to convert")
		}
		pair2High, err := strconv.Atoi(second[1])
		if err != nil {
			log.Fatal("Invalid string to convert")
		}
		pair1 := Pair{
			lowBound:  pair1Low,
			highBound: pair1High,
		}
		pair2 := Pair{
			lowBound:  pair2Low,
			highBound: pair2High,
		}
		if isPairInside(pair1, pair2) {
			count++
			fmt.Printf("%v inside, %v, %v \n", count, pair1, pair2)
		}
	}
}

func isPairInside(pair1 Pair, pair2 Pair) bool {
	if (pair1.lowBound >= pair2.lowBound && pair1.highBound <= pair2.highBound) ||
		(pair2.lowBound >= pair1.lowBound && pair2.highBound <= pair1.highBound) {
		return true
	}

	return false
}
