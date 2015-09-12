package main

import (
	"github.com/twoodhouse/coup-sim/controller"
	"strconv"
	"github.com/twoodhouse/coup-sim/model/strategy"
	"fmt"
)

func main() {
	var numPlayers int
	fmt.Printf("Hello, world. How many players will there be today?\n")
	fmt.Scanf("%d", &numPlayers)
	fmt.Printf("Ah ... %d. Sounds good. What strategies will they be using?\n", numPlayers)

	strategies := make([]*strategy.Interface, numPlayers)
	playerNames := make([]string, numPlayers)

	var strategyName string
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("> ")
		ni, err := fmt.Scanf("%s", &strategyName)
		if err != nil {
				fmt.Println(ni, err)
				return
		}
		newStrategy := createStrategyByName(strategyName)
		strategies[i] = &newStrategy
		if newStrategy != nil {
			fmt.Printf("Found %q, what is its name?  ", newStrategy.GetStrategyName())
			//TODO fix this so that it will error check for multiple players with same name
			ni, err := fmt.Scanf("%s", &playerNames[i])
			if err != nil {
					fmt.Println(ni, err)
					return
			}
			newStrategy.SetPlayerName(playerNames[i])
			if numPlayers - i != 1 {
				fmt.Println("What is the strategy for the next player?")
			}
		} else {
			fmt.Printf("Strategy %q\n does not exist. Try again.\n", strategyName)
			i = i - 1
		}
	}
	var numGames int
	fmt.Printf("How many times do you want to play?\n")
	fmt.Scanf("%d", &numGames)

	winMap := make(map[string]int, len(playerNames))
	for i := 0; i<numGames; i++ {
		winner, returnedStrategies := controller.StartGame(strategies, playerNames, numGames)
		strategies = returnedStrategies
		winMap[winner] = winMap[winner] + 1
	}
	for i := range playerNames {
		fmt.Println(playerNames[i] + " - " + strconv.Itoa(winMap[playerNames[i]]) + " wins")
	}
}
