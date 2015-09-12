package main
import (
  "github.com/twoodhouse/coup-sim/model/strategy"
  "github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
  "github.com/twoodhouse/coup-sim/model/strategies/thiefStrategy"
  "github.com/twoodhouse/coup-sim/model/strategies/consoleStrategy"
  "github.com/twoodhouse/coup-sim/model/strategies/randomStrategy"
)

func createStrategyByName(name string) strategy.Interface {
	switch name {
	case "noLie":
		noLieStrategy := noLieStrategy.New()
		return noLieStrategy
	case "thief":
		thiefStrategy := thiefStrategy.New()
		return thiefStrategy
	case "console":
		consoleStrategy := consoleStrategy.New()
		return consoleStrategy
	case "random":
		randomStrategy := randomStrategy.New()
		return randomStrategy
	}
	return nil
}
