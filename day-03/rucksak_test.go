package main

import (
	"fmt"
	"testing"
)

func TestFindSharedItem(t *testing.T) {
	testCases := []struct {
		line       string
		sharedItem string
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "p"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L"},
		{"PmmdzqPrVvPwwTWBwg", "P"},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v"},
		{"ttgJtRGJQctTZtZT", "t"},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", "s"},
	}

	for _, tt := range testCases {
		t.Run(tt.line, func(t *testing.T) {
			sharedItem, err := FindSharedItem(tt.line)
			if err != nil {
				t.Fatalf("FindSharedItem(%q) returned an error: %v", tt.line, err)
			}

			if sharedItem != tt.sharedItem {
				t.Fatalf("FindSharedItem(%q) = %q, want %q", tt.line, sharedItem, tt.sharedItem)
			}
		})
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		char  string
		score int
	}{
		{"a", 1},
		{"A", 27},
		{"p", 16},
		{"L", 38},
		{"P", 42},
		{"v", 22},
		{"t", 20},
		{"s", 19},
	}

	for _, tt := range tests {
		t.Run(tt.char, func(t *testing.T) {
			score := Score(tt.char)
			if score != tt.score {
				t.Fatalf("Score(%q) = %d, want %d", tt.char, score, tt.score)
			}
		})
	}
}

func TestFindOverlap(t *testing.T) {
	tets := []struct {
		lines   []string
		overlap string
	}{
		{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, "r"},
		{[]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}, "Z"},
	}

	for i, tt := range tets {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			overlap, err := FindOverlap(tt.lines)
			if err != nil {
				t.Fatalf("FindOverlap(%q) returned an error: %v", tt.lines, err)
			}

			if overlap != tt.overlap {
				t.Fatalf("FindOverlap(%q) = %q, want %q", tt.lines, overlap, tt.overlap)
			}
		})
	}
}
