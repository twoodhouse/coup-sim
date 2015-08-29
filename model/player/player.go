package player
import "github.com/twoodhouse/coup-sim/model/deck"
import "github.com/twoodhouse/coup-sim/model/strategy"

type Entity struct {
  name string
  strategy strategy.Interface
  deck deck.Entity
  coins int
}

func New(name string, strategy strategy.Interface, deck deck.Entity, coins int) Entity {
  var entity = Entity {
    name,
    strategy,
    deck,
    coins,
  }
  return entity
}

func (entity *Entity) addCoins(number int) {
  entity.coins = entity.coins + number;
}

func (entity *Entity) Name() string {
  return entity.name
}
