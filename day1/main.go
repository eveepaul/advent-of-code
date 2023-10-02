package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var highestElf Elf
	var tempElf Elf
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			if tempElf.Cal > highestElf.Cal {
				highestElf = tempElf
			}
			tempElf.ID = tempElf.ID + 1
			tempElf.Cal = 0
			continue
		}
		calo, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Invalid int: %s \n", scanner.Text())
		}
		tempElf.Cal = tempElf.Cal + calo
		fmt.Printf("Elf: %v, Calo: %v, Total calo %v \n", tempElf.ID, scanner.Text(), tempElf.Cal)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanner")
	}

	fmt.Printf("Elf: %v is carrying the highest calories. (%v)", highestElf.ID, highestElf.Cal)
}
