package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day-03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		char, err := FindSharedItem(line)
		if err != nil {
			log.Fatalf("Error finding shared item: %s: %v", line, err)
		}

		totalScore += Score(char)
	}

	log.Printf("Total score: %d", totalScore)
}
