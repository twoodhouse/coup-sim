package controller

import (
	"fmt"
	"github.com/twoodhouse/coup-sim/model/strategy"
	"github.com/twoodhouse/coup-sim/model/log"
  "github.com/twoodhouse/coup-sim/model/table"
	// "github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
)

func StartGame(strategies []strategy.Interface, numGames int) {
  table := table.New(strategies)
  players := table.Players()
  log := log.New()
  gameComplete := false
  for turnCounter := 1; !gameComplete; turnCounter++ {
    currentPlayer := players[turnCounter % len(players)]

    action := currentPlayer.Strategy().GetAction(log, table.PlayerCoins(), table.FaceupDecks())
    log.SetPlayerName(currentPlayer.Name())
    log.SetActionName(action)


    if turnCounter == 7 {
      gameComplete = true
    }
  }

  fmt.Println(log.PrettyJsonStr())
}
