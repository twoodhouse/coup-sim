package table
import (
  "github.com/twoodhouse/coup-sim/model/player"
  "github.com/twoodhouse/coup-sim/model/strategy"
  "github.com/twoodhouse/coup-sim/model/deck"
)

type Entity struct {
  players []player.Entity
  centerDeck deck.Entity
  faceupDecks []deck.Entity
  coins int
}

func New(strategies []strategy.Interface) Entity {
  centerDeck := deck.NewRandomCenter()
  faceupDecks := make([]deck.Entity, len(strategies))

  players := make([]player.Entity, 6)
  for i := 0; i < len(strategies); i++ {
    players[i] = player.New(strategies[i].GetName(), strategies[i], deck.New(centerDeck.TakeCards(2)), 2)
  }

  for i := 0; i < len(faceupDecks); i++ {
    faceupDecks[i] = deck.New(make([]int, 2))
  }
  var entity = Entity {
    players,
    centerDeck,
    faceupDecks,
    40,
  }
  return entity
}

func (entity *Entity) FaceupDecks() []deck.Entity {
  return entity.faceupDecks
}

func (entity *Entity) AddCoins(number int) {
  entity.coins = entity.coins + number;
}
