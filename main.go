package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: rank PATH_TO_GAME_RESULTS")
		os.Exit(1)
	}

	gameResults, err := loadGameResults(os.Args[1])
	if err != nil {
		log.Fatalln("failed to load game results: ", err)
	}

	ranking, err := calculateRanking(gameResults)
	if err != nil {
		log.Fatalln("failed to calculate ranking: ", err)
	}

	printRanking(ranking)
}
