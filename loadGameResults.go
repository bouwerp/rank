package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func loadGameResults(path string) ([]GameResult, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}

	data, err := io.ReadAll(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	gameResults, err := parseGameResults(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input file: %w", err)
	}

	return gameResults, nil
}

// parseGameResults parses the body of a game results file and returns a slice of game results.
func parseGameResults(data []byte) ([]GameResult, error) {
	// replace windows line endings
	dataSanitized := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(dataSanitized, "\n")
	gameResults := make([]GameResult, 0)
	for idx, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			// skip empty lines
			continue
		}
		sides := strings.Split(line, ",")
		if len(sides) != 2 {
			return nil, fmt.Errorf("malformed game definition on line %d", idx)
		}
		sideA, err := parseSide(sides[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse side A on line %d: %w", idx, err)
		}
		sideB, err := parseSide(sides[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse side B on line %d: %w", idx, err)
		}
		gameResults = append(gameResults, GameResult{
			SideA: *sideA,
			SideB: *sideB,
		})
	}
	return gameResults, nil
}

// parseSide parses the string representation of a side in a game, expecting the format `TEAM_NAME, SCORE`
// (without the quotes), where score is an integer.
func parseSide(sideStr string) (*Side, error) {
	parts := strings.Split(strings.TrimSpace(sideStr), " ")
	if len(parts) < 2 {
		return nil, errors.New("malformed side")
	}
	score, err := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
	if err != nil {
		return nil, fmt.Errorf("failed to parse score: %w", err)
	}
	return &Side{
		Name:  strings.TrimSpace(strings.Join(parts[:len(parts)-1], " ")),
		Score: score,
	}, nil
}
