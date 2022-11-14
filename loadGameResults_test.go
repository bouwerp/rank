package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_loadGameResults(t *testing.T) {
	type TestCase struct {
		description string
		path        string
		check       func(*testing.T, []GameResult, error)
	}

	testCases := []TestCase{
		{
			description: "positive",
			path:        "test_input",
			check: func(t *testing.T, results []GameResult, err error) {
				assert.Nil(t, err)
				assert.Len(t, results, 5)
				for idx, result := range results {
					switch idx {
					case 0:
						assert.Equal(t, GameResult{
							SideA: Side{
								Name:  "Lions",
								Score: 3,
							},
							SideB: Side{
								Name:  "Snakes",
								Score: 3,
							},
						}, result)
					case 1:
						assert.Equal(t, GameResult{
							SideA: Side{
								Name:  "Tarantulas",
								Score: 1,
							},
							SideB: Side{
								Name:  "FC Awesome",
								Score: 0,
							},
						}, result)
					case 2:
						assert.Equal(t, GameResult{
							SideA: Side{
								Name:  "Lions",
								Score: 1,
							},
							SideB: Side{
								Name:  "FC Awesome",
								Score: 1,
							},
						}, result)
					case 3:
						assert.Equal(t, GameResult{
							SideA: Side{
								Name:  "Tarantulas",
								Score: 3,
							},
							SideB: Side{
								Name:  "Snakes",
								Score: 1,
							},
						}, result)
					case 4:
						assert.Equal(t, GameResult{
							SideA: Side{
								Name:  "Lions",
								Score: 4,
							},
							SideB: Side{
								Name:  "Grouches",
								Score: 0,
							},
						}, result)
					}
				}
			},
		},
	}

	for _, tc := range testCases {
		gameResults, err := loadGameResults(tc.path)
		tc.check(t, gameResults, err)
	}
}
