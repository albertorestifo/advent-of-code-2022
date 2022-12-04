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

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		char, err := FindSharedItem(line)
		if err != nil {
			log.Fatalf("Error finding shared item: %s: %v", line, err)
		}

		totalScore += Score(char)
	}

	var priorities int
	for _, group := range chunks(lines, 3) {
		overlap, err := FindOverlap(group)
		if err != nil {
			log.Fatalf("Error finding overlap: %v", err)
		}

		priorities += Score(overlap)
	}

	log.Printf("Total score: %d", totalScore)
	log.Printf("Total priorities: %d", priorities)
}

func chunks[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
