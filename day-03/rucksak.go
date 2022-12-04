package main

import (
	"errors"
	"strings"
)

func FindSharedItem(line string) (string, error) {
	chars := strings.Split(line, "")

	seenChar := make(map[string]bool)
	midPoint := len(chars) / 2

	for i, char := range chars {
		if i < midPoint {
			// Populate the seen map
			seenChar[char] = true
			continue
		}

		// Try to see if this character has already been seen.
		// If that's the case, then we have found the shared item.
		if seenChar[char] {
			return char, nil
		}
	}

	return "", errors.New("No shared item found")
}

func Score(char string) int {
	ascii := int(rune(char[0]))

	if ascii >= 97 {
		return ascii - 96
	}

	return ascii - 38
}

// Find the character common among the 3 passed lines
func FindOverlap(lines []string) (string, error) {
	seenLine := make(map[string]int)

	for _, line := range lines {
		seenChar := make(map[string]bool)
		for _, char := range strings.Split(line, "") {
			seenChar[char] = true
		}

		for char, _ := range seenChar {
			seenLine[char]++
		}
	}

	// Find the only char that was seen exactly len(lines) times
	for char, seen := range seenLine {
		if seen == len(lines) {
			return char, nil
		}
	}

	return "", errors.New("no overlap found")
}
