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
    otherPlayers := make([]*player.Entity, len(players) - 1)

    //populate otherPlayers
    addedCurrentPlayerCorrection := 0
    for i := 0 ; i < len(players); i++ {
      if i != turnCounter % len(players) {
        otherPlayers[i - addedCurrentPlayerCorrection] = players[i]
      } else {
        addedCurrentPlayerCorrection = 1
      }
    }

    log, table = DoTurn(currentPlayer, otherPlayers, log, table, turnCounter)
    if turnCounter == 3 {
      gameComplete = true
    } else {
      log.NextTurn()
    }
  }

  fmt.Println(log.PrettyJsonStr())
	table.PrintTable()
}

func DoTurn(player *player.Entity, otherPlayers []*player.Entity, log *log.Entity, table *table.Entity, turnCounter int) (*log.Entity, *table.Entity) {
  action := player.Strategy().GetAction(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck())
  log.SetPlayerName(player.Name())
  log.SetActionName(action)
	// playerDeck := player.Deck()
	// var card int

	//top-level challenge logic - it can only happen with these actions
	if action == "tax" || action == "steal" || action == "assassinate" || action == "exchange" {

		for i := 0; i < len(otherPlayers); i++ {
			if otherPlayers[i].Strategy().GetChallenge(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
				challengeSuccess := !player.Deck().HasCardForAction(action)

				if challengeSuccess {
					lossChoice := player.Strategy().GetLossChoice(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck())
					fmt.Println(lossChoice)
					// fmt.Println(playerDeck.Size())
					card := player.Deck().TakeTopCard()
					playerFaceupDeck := table.FaceupDecks()[player.Name()]
					if playerFaceupDeck[0] != 0 {
						playerFaceupDeck[1] = playerFaceupDeck[0]
					}
					playerFaceupDeck[0] = card
				}

				log.CreateChallenge(otherPlayers[i].Name(), challengeSuccess)
			}
		}

	}

  // if action == "steal" || action == "foreign aid" || action == "assassinate" {
	//
  // }

  return log, table
}
