package main

import (
	"fmt"
	"strings"
)

// printRanking formats the provided ranking and prints the result
func printRanking(ranking []RankEntry) {
	outputLines := make([]string, 0)
	for _, rankEntry := range ranking {
		outputLines = append(outputLines,
			fmt.Sprintf("%d. %s, %d pts", rankEntry.Rank, rankEntry.Name, rankEntry.Points))
	}
	output := strings.Join(outputLines, "\n")
	fmt.Println(output)
}
