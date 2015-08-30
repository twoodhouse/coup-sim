package noLieStrategy

import "github.com/twoodhouse/coup-sim/model/log"

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

func (entity *Entity) GetAction(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int) string {
  return "income"
}

func (entity *Entity) GetChallenge(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int) bool {
  return false
}

func (entity *Entity) GetBlock(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int) bool {
  return false
}

func (entity *Entity) GetAmbassadorReturns(log log.Entity, coinInfo map[string]int, faceupInfo map[string][]int, c1 string, c2 string) (string, string) {
  return c1, c2
}
