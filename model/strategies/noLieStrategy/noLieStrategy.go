package noLieStrategy

import "github.com/twoodhouse/coup-sim/model/log"
import "github.com/twoodhouse/coup-sim/model/deck"
// import "fmt"

type Entity struct {
  strategyName string
  playerName string
}

func New() *Entity {
  var entity = Entity {
    "noLieStr",
    "notSet",
  }
  return &entity
}

func (entity *Entity) GetStrategyName() string {
  return entity.strategyName
}

func (entity *Entity) GetDuelFirstCardChoice() int {
  return 1
}

func (entity *Entity) SetPlayerName(name string) {
  entity.playerName = name
}

func (entity *Entity) GetAction(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) string {
  return "tax"
}

func (entity *Entity) GetLossChoice(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) int {
  return 0
}

func (entity *Entity) GetTarget(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) string {
  return "a"
}

func (entity *Entity) GetChallenge(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) bool {
  return true
}

func (entity *Entity) GetBlock(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) bool {
  return false
}

func (entity *Entity) GetStealBlockCardChoice(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) int {
  return 2
}

func (entity *Entity) GetAmbassadorReturns(log *log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, c1 int, c2 int, deck *deck.Entity) (int, int) {
  return c1, c2
}
