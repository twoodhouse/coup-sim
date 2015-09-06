package controller

import (
	"fmt"
	"github.com/twoodhouse/coup-sim/model/strategy"
	"github.com/twoodhouse/coup-sim/model/log"
  "github.com/twoodhouse/coup-sim/model/table"
  "github.com/twoodhouse/coup-sim/model/player"
	"strconv"
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

//order of turn operation:
//1. Assign action and target
//2. Check general challenge concerning action
//3. Check block from targeted player
//4. Check challenge concerning block from source player
//5. Set action results given challenge/block success
func DoTurn(player *player.Entity, otherPlayers []*player.Entity, log *log.Entity, table *table.Entity, turnCounter int) (*log.Entity, *table.Entity) {
  action := player.Strategy().GetAction(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
	targetedPlayer := player //set default - not used unless assassinate or steal happens
  log.SetPlayerName(player.Name())
  log.SetActionName(action)
	challengeLoss := false
	var target string

	//mid-game coin outputs
	// fmt.Print(player.Name() + ": ")
	// fmt.Print(player.Coins())
	// fmt.Println()

	if action != "coup" && player.Coins() >= 10{
		disqualifyPlayer(player, table, log, "Player had 10+ coins and did not coup")
	}

	//get target for certain actions
	if action == "steal" || action == "assassinate" || action == "coup" && !player.Dead(){
		target = player.Strategy().GetTarget(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
		if validLiveOtherPlayerName(otherPlayers, target) {
			log.CreateTarget(target)
		} else {
			disqualifyPlayer(player, table, log, "Target '" + target + "' is not a valid opposing-player name")
		}
	}

	//top-level challenge logic - it can only happen with these actions
	if (action == "tax" || action == "steal" || action == "assassinate" || action == "exchange") && !player.Dead(){
		for i := 0; i < len(otherPlayers); i++ {
			if !otherPlayers[i].Dead() && otherPlayers[i].Strategy().GetChallenge(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), otherPlayers[i].Deck()) {
				challengeSuccess := !player.Deck().HasCardForAction(action)
				losingPlayer := player
				if challengeSuccess {
					challengeLoss = true
					losingPlayer = player
				} else {
					losingPlayer = otherPlayers[i]
				}
				cardLoss := revealCard(losingPlayer, table, log)
				log.CreateChallenge(otherPlayers[i].Name(), challengeSuccess, cardLoss)
				break
			}
		}
	}

	if (action == "steal" || action == "assassinate") && !player.Dead() && !challengeLoss && !targetedPlayer.Dead(){
		if targetedPlayer.Strategy().GetBlock(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
			log.CreateBlock()
			var blockCardClaim int
			if action == "steal" {
				blockCardClaim = targetedPlayer.Strategy().GetStealBlockCardChoice(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
				log.CreateBlockCardClaim(blockCardClaim)
			}
			if player.Strategy().GetChallenge(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
				var neededCardAction string
				if action == "steal" {
					if blockCardClaim == 2 {
						neededCardAction = "steal"
					} else if blockCardClaim == 5 {
						neededCardAction = "exchange"
					} else {
						disqualifyPlayer(targetedPlayer, table, log, "Block card '" + strconv.Itoa(blockCardClaim) + "' is not valid for blocking steal'")
					}
				} else if action == "assassinate" {
					neededCardAction = "block"
				}
				challengeSuccess := !player.Deck().HasCardForAction(neededCardAction)
				losingPlayer := player
				if challengeSuccess {
					losingPlayer = targetedPlayer
				} else {
					challengeLoss = true
					losingPlayer = player
				}
				cardLoss := revealCard(losingPlayer, table, log)
				log.CreateBlockChallenge(challengeSuccess, cardLoss)
			} else {
				challengeLoss = true
			}
		}
	}

	if action == "foreign aid" && !player.Dead() && !challengeLoss && !targetedPlayer.Dead() {
		for i := 0; i < len(otherPlayers); i++ {
			if !otherPlayers[i].Dead() && otherPlayers[i].Strategy().GetBlock(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
				log.CreateBlock()
				log.CreateBlocker(otherPlayers[i].Name())
				if player.Strategy().GetChallenge(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
					challengeSuccess := !otherPlayers[i].Deck().HasCardForAction("tax")
					losingPlayer := player
					if challengeSuccess {
						challengeLoss = true
						losingPlayer = otherPlayers[i]
					} else {
						losingPlayer = player
					}
					cardLoss := revealCard(losingPlayer, table, log)
					log.CreateBlockChallenge(challengeSuccess, cardLoss)
				}
				break
			}
		}
	}

	if (action == "tax" || action == "steal" || action == "coup") && !player.Dead(){
		targetedPlayer = playerByName(otherPlayers, target)
	}

	if action == "tax" && !challengeLoss && !player.Dead(){
		player.AddCoins(3)
		table.AddCoins(-3)
	}

	if action == "steal" && !challengeLoss && !player.Dead(){
		player.AddCoins(2)
		targetedPlayer.AddCoins(-2)
	}

	if action == "coup" && !player.Dead(){
		player.AddCoins(-7)
		log.CreateCardKilled(revealCard(targetedPlayer, table, log))
	}

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

func validLiveOtherPlayerName(otherPlayers []*player.Entity, name string) bool {
	for i := 0; i < len(otherPlayers); i++ {
		if otherPlayers[i].Name() == name && !otherPlayers[i].Dead(){
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

func revealCard(player *player.Entity, table *table.Entity, log *log.Entity) int {
	var cardLoss int
	lossChoice := player.Strategy().GetLossChoice(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
	if lossChoice == 0 {
		cardLoss = player.Deck().TakeTopCard()
	} else {
		cardLoss = player.Deck().TakeBottomCard()
	}
	losingPlayerFaceupDeck := table.FaceupDecks()[player.Name()]
	if losingPlayerFaceupDeck[0] != 0 {
		losingPlayerFaceupDeck[1] = losingPlayerFaceupDeck[0]
	}
	losingPlayerFaceupDeck[0] = cardLoss
	return cardLoss
}
