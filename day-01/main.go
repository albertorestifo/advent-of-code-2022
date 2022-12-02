package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./day-01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elfTotal := make([]int, 0)
	currentElf := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calories := scanner.Text()
		if calories == "" {
			currentElf++
			continue
		}

		calInt, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatalf("Error converting calories %s to int: %v", calories, err)
		}

		if len(elfTotal) < currentElf+1 {
			elfTotal = append(elfTotal, 0)
		}

		elfTotal[currentElf] += calInt
	}

	// Sort the list by calories
	sort.Sort(sort.Reverse(sort.IntSlice(elfTotal)))

	// Find the elf with the most calories
	log.Printf("The elf with the most calories is %d", elfTotal[0])

	// Top 3 elves summed up:
	log.Printf("The top 3 elves have %d calories", elfTotal[0]+elfTotal[1]+elfTotal[2])
}
