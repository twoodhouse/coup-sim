package noLieStrategy
import "fmt"

type Entity struct {
  name string
}

func New() Entity {
  var entity = Entity {
    "testname",
  }
  return entity
}

func (entity *Entity) GetAction() {
  fmt.Printf("Interface works!")
}
