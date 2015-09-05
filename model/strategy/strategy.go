package strategy
import "github.com/twoodhouse/coup-sim/model/log"
import "github.com/twoodhouse/coup-sim/model/deck"

type Interface interface {
  //sets your game player name for internal use
  SetPlayerName(string)

  //return an arbitrary strategy-type name for interface use
  GetStrategyName() string

  //return int (1-5)
  GetDuelFirstCardChoice() int

  //return one of seven action names, {"income", "foreign aid", "tax", "steal", "assassinate", "exchange", "coup"}
  //if player has more than 10 coins, player must return "coup" or will be disqualified
  GetAction(*log.Entity, map[string]int, map[string][]int, *deck.Entity) string

  //return a valud player
  GetTarget(*log.Entity, map[string]int, map[string][]int, *deck.Entity) string

  //return 0 or 1 to specify which of your cards will be lost
  GetLossChoice(*log.Entity, map[string]int, map[string][]int, *deck.Entity) int

  //return challenge choice for the current player's action
  GetChallenge(*log.Entity, map[string]int, map[string][]int, *deck.Entity) bool

  //return block choice since you have been targeted by a player
  GetBlock(*log.Entity, map[string]int, map[string][]int, *deck.Entity) bool

  //return a 2 or a 5 (captain or ambassador respectively)
  GetStealBlockCardChoice(*log.Entity, map[string]int, map[string][]int, *deck.Entity) int

  //return two ints (1-5) which are the returned cards
  GetAmbassadorReturns(*log.Entity, map[string]int, map[string][]int, int, int, *deck.Entity) (int,int)
}
