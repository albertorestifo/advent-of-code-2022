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
