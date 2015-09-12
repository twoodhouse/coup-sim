package controller

import (
	// "fmt"
	"github.com/twoodhouse/coup-sim/model/strategy"
	"github.com/twoodhouse/coup-sim/model/log"
  "github.com/twoodhouse/coup-sim/model/table"
  "github.com/twoodhouse/coup-sim/model/player"
	"strconv"
)

func StartGame(strategies []*strategy.Interface, playerNames []string, numGames int) (string,[]*strategy.Interface) {
  table := table.New(strategies, playerNames)
  players := table.Players()
  log := log.New()
  gameComplete := false
	var winner string
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
			for i := range players {
				if players[i].Deck().Size() > 0 {
					winner = players[i].Name()
				}
			}
    } else {
			if !deadTurn {
				log.NextTurn()
			}
    }
  }
  // fmt.Println(log.PrettyJsonStr())
	// table.PrintTable()
	return winner, strategies
}

//order of turn operation:
//1. Assign action and target
//2. Check general challenge concerning action
//3. Check block from targeted player
//4. Check challenge concerning block from source player
//5. Set action results given challenge/block success
func DoTurn(player *player.Entity, otherPlayers []*player.Entity, log *log.Entity, table *table.Entity, turnCounter int) (*log.Entity, *table.Entity) {
	log.SetPlayerName(player.Name())
	action := player.Strategy().GetAction(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
	targetedPlayer := player //set default - not used unless assassinate or steal happens
  log.SetActionName(action)
	challengeLoss := false
	var target string

	//mid-game coin outputs
	// fmt.Print(player.Name() + ": ")
	// fmt.Print(player.Coins())
	// fmt.Println()

	if !player.Dead() && action != "tax" && action != "income" && action != "foreign_aid" && action != "coup" && action != "steal" && action != "assassinate" && action != "exchange" && action != "coup" {
		disqualifyPlayer(player, table, log, "Player attempted invalid action " + action)
	}

	if action != "coup" && player.Coins() >= 10 && !player.Dead() {
		disqualifyPlayer(player, table, log, "Player had 10+ coins and did not coup")
	}

	if action == "coup" && player.Coins() < 7 && !player.Dead() {
		disqualifyPlayer(player, table, log, "Player had " + strconv.Itoa(player.Coins()) + " coins and needed 7 to Coup")
	}

	if action == "assassinate" && !player.Dead() {
		if player.Coins() < 3 {
			disqualifyPlayer(player, table, log, "Player had " + strconv.Itoa(player.Coins()) + " coins and needed 3 to Assassinate")
		} else {
			player.AddCoins(-3)
		}
	}

	//get target for certain actions
	if action == "steal" || action == "assassinate" || action == "coup" && !player.Dead(){
		target = player.Strategy().GetTarget(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
		if validLiveOtherPlayerName(otherPlayers, target) {
			targetedPlayer = playerByName(otherPlayers, target)
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
					var cardValue int
					switch action {
					case "tax":
						cardValue = 1
					case "steal":
						cardValue = 2
					case "assassinate":
						cardValue = 3
					case "block":
						cardValue = 4
					case "exchange":
						cardValue = 5
					}
					swapChallengedCard(player, table, log, cardValue)
				}
				log.CreateChallenge(otherPlayers[i].Name(), challengeSuccess)
				cardLoss := revealCard(losingPlayer, table, log)
				log.CreateChallengeCardLoss(cardLoss)
				break
			}
		}
	}

	if (action == "steal" || action == "assassinate") && !player.Dead() && !challengeLoss && !targetedPlayer.Dead(){
		if targetedPlayer.Strategy().GetBlock(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), targetedPlayer.Deck()) {
			log.CreateBlock()
			var blockCardClaim int
			if action == "steal" {
				blockCardClaim = targetedPlayer.Strategy().GetStealBlockCardChoice(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), targetedPlayer.Deck())
				log.CreateBlockCardClaim(blockCardClaim)
				if blockCardClaim != 2 && blockCardClaim != 5 {
					disqualifyPlayer(targetedPlayer, table, log, "Block card '" + strconv.Itoa(blockCardClaim) + "' is not valid for blocking steal'")
				}
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
				challengeSuccess := !targetedPlayer.Deck().HasCardForAction(neededCardAction)
				losingPlayer := player
				var cardValue int
				if challengeSuccess {
					losingPlayer = targetedPlayer
				} else {
					challengeLoss = true
					losingPlayer = player
					switch neededCardAction {
					case "tax":
						cardValue = 1
					case "steal":
						cardValue = 2
					case "assassinate":
						cardValue = 3
					case "block":
						cardValue = 4
					case "exchange":
						cardValue = 5
					}
					swapChallengedCard(targetedPlayer, table, log, cardValue)
				}
				log.CreateBlockChallenge(challengeSuccess)
				cardLoss := revealCard(losingPlayer, table, log)
				log.CreateBlockChallengeCardLoss(cardLoss)
			} else {
				challengeLoss = true
			}
		}
	}

	if action == "foreign_aid" && !player.Dead() && !challengeLoss && !targetedPlayer.Dead() {
		for i := 0; i < len(otherPlayers); i++ {
			if !otherPlayers[i].Dead() && otherPlayers[i].Strategy().GetBlock(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), otherPlayers[i].Deck()) {
				log.CreateBlock()
				challengeLoss = true
				log.CreateBlocker(otherPlayers[i].Name())
				if player.Strategy().GetChallenge(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck()) {
					challengeSuccess := !otherPlayers[i].Deck().HasCardForAction("tax")
					losingPlayer := player
					if challengeSuccess {
						losingPlayer = otherPlayers[i]
						challengeLoss = false
					} else {
						losingPlayer = player
						swapChallengedCard(otherPlayers[i], table, log, 1)
					}
					log.CreateBlockChallenge(challengeSuccess)
					cardLoss := revealCard(losingPlayer, table, log)
					log.CreateBlockChallengeCardLoss(cardLoss)
				}
				break
			}
		}
	}

	if action == "income" {
		player.AddCoins(1)
		table.AddCoins(-1)
	}

	if action == "foreign_aid" && !challengeLoss && !player.Dead(){
		player.AddCoins(2)
		table.AddCoins(-2)
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

	if action == "assassinate" && !challengeLoss && !player.Dead() && !targetedPlayer.Dead(){
		log.CreateCardKilled(revealCard(targetedPlayer, table, log))
	}

	if action == "exchange" && !challengeLoss && !player.Dead() {
		table.CenterDeck().ShuffleCards()
		player.Deck().ShuffleCards()
		c1 := table.CenterDeck().TakeCards(1)[0]
		c2 := table.CenterDeck().TakeCards(1)[0]
		player.Deck().GiveCards([]int{c1,c2})
		r1,r2 := player.Strategy().GetExchangeReturnChoices(log, table.PlayerNames(), table.PlayerCoins(), table.FaceupDecks(), player.Deck())
		if !player.Deck().HasCards(r1, r2) {
			disqualifyPlayer(player, table, log, "Player gave invalid card returns " + strconv.Itoa(r1) + "," + strconv.Itoa(r2))
			table.CenterDeck().GiveCards([]int{c1,c2})
		} else {
			fourCards := player.Deck().TakeCards(player.Deck().Size())
			returnToTableCards := make([]int, 2)
			r1found := false
			r2found := false
			for i := range fourCards {
				var index int
				if returnToTableCards[1] == 0 {
					index = 1
				}
				if returnToTableCards[0] == 0 {
					index = 0
				}
				if fourCards[i] == r1 && !r1found{
					r1found = true
					returnToTableCards[index] = fourCards[i]
				} else if fourCards[i] == r2 && !r2found{
					r2found = true
					returnToTableCards[index] = fourCards[i]
				} else {
					player.Deck().GiveCards([]int{fourCards[i]})
				}
			}
			table.CenterDeck().GiveCards(returnToTableCards)
			player.Deck().ShuffleCards()
		}
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
		cardLoss = player.Deck().TakeBottomCard()
	} else if lossChoice == 1 {
		cardLoss = player.Deck().TakeTopCard()
	} else {
		disqualifyPlayer(player, table, log, "invalid logChoice value" + strconv.Itoa(lossChoice))
		cardLoss = 0
	}
	if !player.Dead() {
		losingPlayerFaceupDeck := table.FaceupDecks()[player.Name()]
		if losingPlayerFaceupDeck[0] != 0 {
			losingPlayerFaceupDeck[1] = losingPlayerFaceupDeck[0]
		}
		losingPlayerFaceupDeck[0] = cardLoss
		if table.FaceupDecks()[player.Name()][0] != 0 && table.FaceupDecks()[player.Name()][1] != 0 {
			player.Kill()
		}
	}
	return cardLoss
}

func swapChallengedCard(player *player.Entity, table *table.Entity, log *log.Entity, card int) {
	table.CenterDeck().GiveCards([]int{card})
	table.CenterDeck().ShuffleCards()
	player.Deck().ReplaceCardWithCard(card, table.CenterDeck().TakeCards(1)[0])
}
