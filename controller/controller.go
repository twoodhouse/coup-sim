package controller

import (
	"fmt"
	"github.com/twoodhouse/coup-sim/model/strategy"
	"github.com/twoodhouse/coup-sim/model/log"
  "github.com/twoodhouse/coup-sim/model/table"
  "github.com/twoodhouse/coup-sim/model/player"
	// "github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
)

func StartGame(strategies []strategy.Interface, playerNames []string, numGames int) {
  table := table.New(strategies, playerNames)
  players := table.Players()
  log := log.New()
  gameComplete := false
  for turnCounter := 0; !gameComplete; turnCounter++ {
    currentPlayer := players[turnCounter % len(players)]
    otherPlayers := make([]player.Entity, len(players) - 1)

    //populate otherPlayers
    addedCurrentPlayerCorrection := 0
    for i := 0 ; i < len(players); i++ {
      if i != turnCounter % len(players) {
        otherPlayers[i - addedCurrentPlayerCorrection] = players[i]
      } else {
        addedCurrentPlayerCorrection = 1
      }
    }

    log, table = DoTurn(currentPlayer, otherPlayers, log, table)
    if turnCounter == 7 {
      gameComplete = true
    } else {
      log.NextTurn()
    }
  }

  fmt.Println(log.PrettyJsonStr())
}

func DoTurn(player player.Entity, otherPlayers []player.Entity, log log.Entity, table table.Entity) (log.Entity, table.Entity) {
  action := player.Strategy().GetAction(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck())
  log.SetPlayerName(player.Name())
  log.SetActionName(action)

	if action == "tax" || action == "steal" || action == "assassinate" || action == "exchange" {
		for i := 0; i < len(otherPlayers); i++ {
			if otherPlayers[i].Strategy().GetChallenge(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
				log.CreateChallenge(otherPlayers[i].Name(), false)
			}
		}
	}

  // if action == "steal" || action == "foreign aid" || action == "assassinate" {
	//
  // }

  return log, table
}
