package table
import "github.com/twoodhouse/coup-sim/model/player"
import "github.com/twoodhouse/coup-sim/model/deck"

type Entity struct {
  players []player.Entity
  centerDeck deck.Entity
  faceupDecks []deck.Entity
  coins int
}

func New(players []player.Entity) Entity {
  centerDeck := deck.NewRandomCenter()
  faceupDecks := make([]deck.Entity, len(players))

  for i := 0; i < len(faceupDecks); i++ {
    faceupDecks[i] = deck.New(make([]int, 2));
  }
  var entity = Entity {
    players,
    centerDeck,
    faceupDecks,
    20,
  }
  return entity
}

func (entity *Entity) Players() []player.Entity {
  return entity.players
}

func (entity *Entity) Coins() int {
  return entity.coins
}

func (entity *Entity) AddCoins(number int) {
  entity.coins = entity.coins + number;
}
// []player.Properties{{"p1"}}
