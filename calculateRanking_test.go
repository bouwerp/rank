package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateRank(t *testing.T) {
	type TestCase struct {
		description string
		input       []GameResult
		check       func(*testing.T, []RankEntry, error)
	}

	testCases := []TestCase{
		{
			description: "positive - single game, clear win",
			input: []GameResult{{
				SideA: Side{
					Name:  "Team B",
					Score: 2,
				},
				SideB: Side{
					Name:  "Team A",
					Score: 0,
				},
			}},
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.Nil(t, err)
				assert.Len(t, entries, 2)
				for idx, entry := range entries {
					switch idx {
					case 0:
						assert.Equal(t, "Team B", entry.Name)
						assert.Equal(t, 3, entry.Points)
						assert.Equal(t, 1, entry.Rank)
					case 1:
						assert.Equal(t, "Team A", entry.Name)
						assert.Equal(t, 0, entry.Points)
						assert.Equal(t, 2, entry.Rank)
					}
				}
			},
		},
		{
			description: "positive - multiple games, two teams, clear win",
			input: []GameResult{
				{
					SideA: Side{
						Name:  "Team B",
						Score: 2,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 0,
					},
				},
				{
					SideA: Side{
						Name:  "Team B",
						Score: 1,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 1,
					},
				},
				{
					SideA: Side{
						Name:  "Team B",
						Score: 2,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 3,
					},
				},
				{
					SideA: Side{
						Name:  "Team B",
						Score: 3,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 2,
					},
				},
			},
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.Nil(t, err)
				assert.Len(t, entries, 2)
				for idx, entry := range entries {
					switch idx {
					case 0:
						assert.Equal(t, "Team B", entry.Name)
						assert.Equal(t, 7, entry.Points)
						assert.Equal(t, 1, entry.Rank)
					case 1:
						assert.Equal(t, "Team A", entry.Name)
						assert.Equal(t, 4, entry.Points)
						assert.Equal(t, 2, entry.Rank)
					}
				}
			},
		},
		{
			description: "positive - multiple games, multiple teams",
			input: []GameResult{
				{
					SideA: Side{
						Name:  "Team B",
						Score: 2,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 0,
					},
				},
				{
					SideA: Side{
						Name:  "Team D",
						Score: 3,
					},
					SideB: Side{
						Name:  "Team C",
						Score: 1,
					},
				},
				{
					SideA: Side{
						Name:  "Team D",
						Score: 2,
					},
					SideB: Side{
						Name:  "Team B",
						Score: 3,
					},
				},
				{
					SideA: Side{
						Name:  "Team C",
						Score: 3,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 4,
					},
				},
			},
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.Nil(t, err)
				assert.Len(t, entries, 4)
				for idx, entry := range entries {
					switch idx {
					case 0:
						assert.Equal(t, "Team B", entry.Name)
						assert.Equal(t, 6, entry.Points)
						assert.Equal(t, 1, entry.Rank)
					case 1:
						assert.Equal(t, "Team A", entry.Name)
						assert.Equal(t, 3, entry.Points)
						assert.Equal(t, 2, entry.Rank)
					case 2:
						assert.Equal(t, "Team D", entry.Name)
						assert.Equal(t, 3, entry.Points)
						assert.Equal(t, 2, entry.Rank)
					case 3:
						assert.Equal(t, "Team C", entry.Name)
						assert.Equal(t, 0, entry.Points)
						assert.Equal(t, 3, entry.Rank)
					}
				}
			},
		},
		{
			description: "negative - empty team name",
			input: []GameResult{
				{
					SideA: Side{
						Name:  "",
						Score: 2,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 0,
					},
				},
			},
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.ErrorContains(t, err, "side A in game 0 has an empty name")
			},
		},
		{
			description: "negative - score < 0",
			input: []GameResult{
				{
					SideA: Side{
						Name:  "Team B",
						Score: -2,
					},
					SideB: Side{
						Name:  "Team A",
						Score: 0,
					},
				},
			},
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.ErrorContains(t, err, "side A in game 0 has a score < 0")
			},
		},
		{
			description: "negative - no games provided (empty)",
			input:       []GameResult{},
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.ErrorContains(t, err, "no games provided")
			},
		},
		{
			description: "negative - no games provided (nil)",
			input:       nil,
			check: func(t *testing.T, entries []RankEntry, err error) {
				assert.ErrorContains(t, err, "no games provided")
			},
		},
	}

	for _, tc := range testCases {
		rank, err := calculateRanking(tc.input)
		tc.check(t, rank, err)
	}
}
