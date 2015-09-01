package noLieStrategy

import "github.com/twoodhouse/coup-sim/model/log"
import "github.com/twoodhouse/coup-sim/model/deck"

type Entity struct {
  name string
}

func New() Entity {
  var entity = Entity {
    "noLieStr",
  }
  return entity
}

func (entity *Entity) GetName() string {
  return entity.name
}

func (entity *Entity) GetAction(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck deck.Entity) string {
  return "tax"
}

func (entity *Entity) GetLossChoice(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck deck.Entity) int {
  return 0
}

func (entity *Entity) GetChallenge(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck deck.Entity) bool {
  return true
}

func (entity *Entity) GetBlock(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, deck deck.Entity) bool {
  return false
}

func (entity *Entity) GetAmbassadorReturns(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, c1 string, c2 string, deck deck.Entity) (string, string) {
  return c1, c2
}
