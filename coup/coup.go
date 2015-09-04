package main

import (
	"github.com/twoodhouse/coup-sim/model/strategy"
	"github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
	"github.com/twoodhouse/coup-sim/model/strategies/thiefStrategy"
	"github.com/twoodhouse/coup-sim/controller"
	"fmt"
)

func main() {
	var numPlayers int
	fmt.Printf("Hello, world. How many players will there be today?\n")
	fmt.Scanf("%d", &numPlayers)
	fmt.Printf("Ah ... %d. Sounds good. What strategies will they be using?\n", numPlayers)

	strategies := make([]strategy.Interface, numPlayers)
	playerNames := make([]string, numPlayers)

	var strategyName string
	for i := 0; i < numPlayers; i++ {
		fmt.Printf(">")
		ni, err := fmt.Scanf("%s", &strategyName)
		if err != nil {
				fmt.Println(ni, err)
				return
		}
		newStrategy := createStrategyByName(strategyName)
		strategies[i] = newStrategy
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

	controller.StartGame(strategies, playerNames, numGames)
}

func createStrategyByName(name string) strategy.Interface {
	switch name {
	case "noLie":
		noLieStrategy := noLieStrategy.New()
		return noLieStrategy
	case "thief":
		thiefStrategy := thiefStrategy.New()
		return thiefStrategy
	}
	return nil
}
