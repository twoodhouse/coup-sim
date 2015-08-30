package main

import (
	// "github.com/twoodhouse/coup-sim/model/strategy"
	"fmt"
)

func main() {
	var numPlayers int
	fmt.Printf("Hello, world. How many players will there be today?\n")
	fmt.Scanf("%d", &numPlayers)
	fmt.Printf("Ah ... %d. Sounds good. What strategies will they be using?\n", numPlayers)
	// strategies := make([]strategy.Interface, numPlayers)
	var strategyName string
	for i := 0; i < numPlayers; i++ {
		fmt.Printf(">")
		ni, err := fmt.Scanf("%s", &strategyName)
		if err != nil {
				fmt.Println(ni, err)
				return
		}
		fmt.Printf("Found it.\n")
		// strategies[i] =
	}
}
