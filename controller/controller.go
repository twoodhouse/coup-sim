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
		playerTurn := turnCounter % len(players)
    currentPlayer := players[playerTurn]
		deadTurn := currentPlayer.Dead()

		otherPlayers := make([]*player.Entity, len(players) - 1)
    for i := 0 ; i < len(players) - 1; i++ {
			if i + playerTurn + 1 < len(players) {
				otherPlayers[i] = players[i + playerTurn + 1]
			} else {
				otherPlayers[i] = players[i + playerTurn + 1 - len(players)]
			}
    }

		if !deadTurn {
			log, table = DoTurn(currentPlayer, otherPlayers, log, table, turnCounter)
		}
    if NumberDead(players) == len(players) - 1 {
      gameComplete = true
    } else {
			if !deadTurn {
				log.NextTurn()
			}
    }
  }

  fmt.Println(log.PrettyJsonStr())
	table.PrintTable()
}

func DoTurn(player *player.Entity, otherPlayers []*player.Entity, log *log.Entity, table *table.Entity, turnCounter int) (*log.Entity, *table.Entity) {
  action := player.Strategy().GetAction(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck())
  log.SetPlayerName(player.Name())
  log.SetActionName(action)
	challengeLoss := false
	var target string

	if action == "steal" || action == "assassinate"{
		target := player.Strategy().GetTarget(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck())
		if validOtherPlayerName(otherPlayers, target) {
			log.CreateTarget(target)
		} else {
			disqualifyPlayer(player, table, log, "Target '" + target + "' is not a valid player name")
		}
	}

	//top-level challenge logic - it can only happen with these actions
	if (action == "tax" || action == "steal" || action == "assassinate" || action == "exchange") && !player.Dead(){
		for i := 0; i < len(otherPlayers); i++ {
			if !otherPlayers[i].Dead() && otherPlayers[i].Strategy().GetChallenge(log, table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
				challengeSuccess := !player.Deck().HasCardForAction(action)
				var cardLoss int
				losingPlayer := player
				if challengeSuccess {
					challengeLoss = true
					losingPlayer = player
				} else {
					losingPlayer = otherPlayers[i]
				}
				lossChoice := losingPlayer.Strategy().GetLossChoice(log, table.PlayerCoins(), table.FaceupDecks(), losingPlayer.Deck())
				if lossChoice == 0 {
					cardLoss = losingPlayer.Deck().TakeTopCard()
				} else {
					cardLoss = losingPlayer.Deck().TakeBottomCard()
				}

				losingPlayerFaceupDeck := table.FaceupDecks()[losingPlayer.Name()]
				if losingPlayerFaceupDeck[0] != 0 {
					losingPlayerFaceupDeck[1] = losingPlayerFaceupDeck[0]
				}
				losingPlayerFaceupDeck[0] = cardLoss

				log.CreateChallenge(otherPlayers[i].Name(), challengeSuccess, cardLoss)
				break
			}
		}
	}

	if action == "tax" && !challengeLoss {
		player.AddCoins(3)
		table.AddCoins(-3)
	}

	if action == "steal" && !challengeLoss {
		player.AddCoins(2)
		playerByName(otherPlayers, target).AddCoins(-2)
	}
  // if action == "steal" || action == "foreign aid" || action == "assassinate" {
	//
  // }
	if table.FaceupDecks()[player.Name()][0] != 0 && table.FaceupDecks()[player.Name()][1] != 0 {
		player.Kill()
	}
	for i := 0; i < len(otherPlayers); i++ {
		if table.FaceupDecks()[otherPlayers[i].Name()][0] != 0 && table.FaceupDecks()[otherPlayers[i].Name()][1] != 0 {
			otherPlayers[i].Kill()
		}
	}

  return log, table
}

func NumberDead(players []*player.Entity) int {
	var num int
	for i := 0; i < len(players); i++ {
		if players[i].Dead() {
			num++
		}
	}
	return num
}

func validOtherPlayerName(otherPlayers []*player.Entity, name string) bool {
	for i := 0; i < len(otherPlayers); i++ {
		if otherPlayers[i].Name() == name {
			return true
		}
	}
	return false
}

func playerByName(players []*player.Entity, name string) *player.Entity {
	for i := 0; i < len(players); i++ {
		if players[i].Name() == name {
			return players[i]
		}
	}
	//not sure to deal with this - I have to return a player, can't do nil
	//make sure to validate the player exists before running this function
	return players[0]
}

func disqualifyPlayer(player *player.Entity, table *table.Entity, log *log.Entity, reason string) {
	player.Kill()
	disqualifiedPlayerFaceupDeck := table.FaceupDecks()[player.Name()]
	if disqualifiedPlayerFaceupDeck[0] != 0 {
		disqualifiedPlayerFaceupDeck[1] = disqualifiedPlayerFaceupDeck[0]
	}
	disqualifiedPlayerFaceupDeck[0] = player.Deck().TakeTopCard();
	if player.Deck().Size() > 0 {
		disqualifiedPlayerFaceupDeck := table.FaceupDecks()[player.Name()]
		if disqualifiedPlayerFaceupDeck[0] != 0 {
			disqualifiedPlayerFaceupDeck[1] = disqualifiedPlayerFaceupDeck[0]
		}
		disqualifiedPlayerFaceupDeck[0] = player.Deck().TakeTopCard();
	}
	log.CreateDisqualify(reason)
}
