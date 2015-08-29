package noLieStrategy
// import "fmt"
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

func (entity *Entity) GetAction(log log.Entity) string {
  return "income"
}

func (entity *Entity) GetChallenge(log log.Entity) bool {
  return false
}

func (entity *Entity) GetBlock(log log.Entity) bool {
  return false
}

func (entity *Entity) GetAmbassadorReturns(log log.Entity, c1 string, c2 string) (string, string) {
  return c1, c2
}
