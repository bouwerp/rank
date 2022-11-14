package main

import (
	"errors"
	"fmt"
	"strings"
)

// calculateRanking calculates an ordered ranking of entries with calculated scores and ranks.
func calculateRanking(games []GameResult) ([]RankEntry, error) {
	if len(games) == 0 {
		return nil, errors.New("no games provided")
	}

	rankEntries := make(map[string]RankEntry)

	for idx, game := range games {
		// validate sides
		if err := validateSide(game.SideA); err != nil {
			return nil, fmt.Errorf("side A in game %d %w", idx, err)
		}
		if err := validateSide(game.SideB); err != nil {
			return nil, fmt.Errorf("side B in game %d %w", idx, err)
		}

		winningSide, losingSide, tie := checkWin(game)

		if tie {
			updateRankEntryPoints(game.SideA, 1, &rankEntries)
			updateRankEntryPoints(game.SideB, 1, &rankEntries)
		} else {
			updateRankEntryPoints(*winningSide, 3, &rankEntries)
			updateRankEntryPoints(*losingSide, 0, &rankEntries)
		}
	}
	return doSortRanking(rankEntries), nil
}

func doSortRanking(rankEntries map[string]RankEntry) []RankEntry {
	return sortRanking(rankEntries, 0, -1)
}

// sortRanking is a recursive function that sorts a provided map of rank entries based on points. Order for tied points
// Rank is further processed for name alphabetical order.
func sortRanking(rankEntries map[string]RankEntry, prevRank, prevPoints int) []RankEntry {
	var first string
	for _, rankEntry := range rankEntries {
		// only for the first iteration when no value has yet been set for first
		if first == "" {
			first = rankEntry.Name
			continue
		}
		// outright higher points means higher rank
		if rankEntry.Points > rankEntries[first].Points {
			first = rankEntry.Name
		}
		// a tie in points and higher alphabetical order means the rank entry is now first in rank
		if rankEntry.Points == rankEntries[first].Points && rankEntry.Name < first {
			first = rankEntry.Name
		}
	}
	firstEntry := rankEntries[first]
	if firstEntry.Points == prevPoints {
		// tied points w.r.t. to previous entry means same rank
		firstEntry.Rank = prevRank
	} else {
		firstEntry.Rank = prevRank + 1
	}
	// remove this first entry
	delete(rankEntries, first)
	if len(rankEntries) > 0 {
		// recursively call this function for remaining entries
		return append([]RankEntry{firstEntry}, sortRanking(rankEntries, firstEntry.Rank, firstEntry.Points)...)
	}
	return []RankEntry{firstEntry}
}

// updateRankEntryPoints checks the existence of the provided side in the collection of rank entries and adds the value
// provided to its current points. If the team is not represented in the rank entries collection it is created with the
// provided points.
func updateRankEntryPoints(side Side, points int, rankEntries *map[string]RankEntry) {
	if rankEntry, exists := (*rankEntries)[side.Name]; exists {
		rankEntry.Points += points
		(*rankEntries)[side.Name] = rankEntry
	} else {
		(*rankEntries)[side.Name] = RankEntry{
			Name:   side.Name,
			Points: points,
		}
	}
}

// checkWin checks which of the two provided sides is the winning side based on their score, and returns pointers to
// the winning and losing sides. If it is a tie, only the boolean indicating it as such is set and returned.
func checkWin(game GameResult) (*Side, *Side, bool) {
	if game.SideA.Score > game.SideB.Score {
		return &game.SideA, &game.SideB, false
	} else if game.SideB.Score > game.SideA.Score {
		return &game.SideB, &game.SideA, false
	} else {
		return nil, nil, true
	}
}

// validateSide performs some common
func validateSide(side Side) error {
	if side.Score < 0 {
		return errors.New("has a score < 0")
	}
	if strings.TrimSpace(side.Name) == "" {
		return errors.New("has an empty name")
	}
	return nil
}
