package main

import (
	"github.com/twoodhouse/coup-sim/model/strategy"
	"github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
	"github.com/twoodhouse/coup-sim/controller"
	"fmt"
)

func main() {
	var numPlayers int
	fmt.Printf("Hello, world. How many players will there be today?\n")
	fmt.Scanf("%d", &numPlayers)
	fmt.Printf("Ah ... %d. Sounds good. What strategies will they be using?\n", numPlayers)

	strategies := make([]strategy.Interface, numPlayers)
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
			fmt.Printf("Found %q\n", newStrategy.GetName())
		} else {
			fmt.Printf("Strategy %q\n does not exist. Try again.\n", strategyName)
			i = i - 1
		}
	}
	var numGames int
	fmt.Printf("How many times do you want to play?\n")
	fmt.Scanf("%d", &numGames)
	
	controller.StartGame(strategies, numGames)
}

func createStrategyByName(name string) strategy.Interface {
	switch name {
	case "noLie":
		noLieStrategy := noLieStrategy.New()
		return &noLieStrategy
	case "other":
		noLieStrategy := noLieStrategy.New()
		return &noLieStrategy
	}
	return nil
}
