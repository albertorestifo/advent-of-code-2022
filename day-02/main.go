package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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

func calculateLineScore(line string) (int, error) {
	elements := strings.Split(line, " ")
	if len(elements) != 2 {
		return 0, fmt.Errorf("line does not have 2 elements: %s", line)
	}

	them, err := mapToConst(elements[0])
	if err != nil {
		return 0, fmt.Errorf("reading them: %v", err)
	}

	us, err := mapToConst(elements[1])
	if err != nil {
		return 0, fmt.Errorf("reading us: %v", err)
	}

	if us == them {
		return DrawScore + us, nil
	}

	if winMatrix[them] == us {
		return 0 + us, nil
	}

	return WinningScore + us, nil
}

func mapToConst(element string) (int, error) {
	switch element {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return 0, fmt.Errorf("unknown element: %s", element)
	}
}
