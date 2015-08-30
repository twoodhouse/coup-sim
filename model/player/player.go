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

func (entity *Entity) AddCoins(number int) {
  entity.coins = entity.coins + number;
}

func (entity *Entity) Name() string {
  return entity.name
}

func (entity *Entity) Coins() int {
  return entity.coins
}

func (entity *Entity) GiveCards(cards []int) {
  entity.deck.GiveCards(cards)
}

func (entity *Entity) DeckSize() int {
  return entity.deck.Size()
}
