package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	_ int = iota
	Rock
	Paper
	Scissors
)

const (
	LostScore    = 0
	DrawScore    = 3
	WinningScore = 6
)

var winMatrix = map[int]int{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var lostMatrix = make(map[int]int)

func main() {
	// Invert the win matrix into the lost matrix
	for key, value := range winMatrix {
		lostMatrix[value] = key
	}

	file, err := os.Open("./day-02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		lineScore, err := calculateLineScore(line)
		if err != nil {
			log.Fatalf("Error calculating line score: %s: %v", line, err)
		}

		totalScore += lineScore
	}

	log.Printf("The total score is %d", totalScore)
}

func calculateLineScore(line string) (int, error) {
	elements := strings.Split(line, " ")
	if len(elements) != 2 {
		return 0, fmt.Errorf("line does not have 2 elements: %s", line)
	}

	them, err := mapToElement(elements[0])
	if err != nil {
		return 0, fmt.Errorf("reading them: %v", err)
	}

	switch elements[1] {
	case "X":
		// We want to loose
		return winMatrix[them], nil

	case "Y":
		// Draw
		return DrawScore + them, nil

	case "Z":
		// We win
		return WinningScore + lostMatrix[them], nil

	default:
		return 0, fmt.Errorf("unknown answer: %s", elements[1])
	}
}

func mapToElement(element string) (int, error) {
	switch element {
	case "A":
		return Rock, nil
	case "B":
		return Paper, nil
	case "C":
		return Scissors, nil
	default:
		return 0, fmt.Errorf("unknown element: %s", element)
	}
}
