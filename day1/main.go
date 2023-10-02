package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	ID  int
	Cal int
}

func main() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var tempElf Elf
	var elves []Elf
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			elves = append(elves, tempElf)
			tempElf.ID = tempElf.ID + 1
			tempElf.Cal = 0
			continue
		}

		calo, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Invalid int: %s \n", scanner.Text())
		}

		tempElf.Cal = tempElf.Cal + calo
		fmt.Printf("Elf: %v, Cal: %v, Total cal %v \n", tempElf.ID, scanner.Text(), tempElf.Cal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanner")
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Cal > elves[j].Cal
	})
	top3Elves := elves[:3]
	fmt.Printf("top 3 elves with highest calories: %v \n", top3Elves)
	var totalTop3Cal int
	for _, elf := range top3Elves {
		totalTop3Cal = totalTop3Cal + elf.Cal
	}

	fmt.Printf("total cal for the top 3 elves: %v", totalTop3Cal)
}
