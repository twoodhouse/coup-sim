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

func New(strategies []strategy.Interface, playerNames []string) Entity {
  centerDeck := deck.NewRandomCenter()
  faceupDecks := make([]deck.Entity, len(strategies))

  players := make([]player.Entity, len(strategies))
  for i := 0; i < len(strategies); i++ {
    players[i] = player.New(playerNames[i], strategies[i], deck.New(centerDeck.TakeCards(2)), 2)
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

func (entity *Entity) Players() []player.Entity {
  return entity.players
}

func (entity *Entity) FaceupDecks() map[string][]int {
  playerDecks := make(map[string][]int)
  for i := 0; i < len(entity.faceupDecks); i++ {
    playerDecks[entity.players[i].Name()] = entity.faceupDecks[i].Cards()
  }
  return playerDecks
}

func (entity *Entity) PlayerCoins() map[string]int {
  playerCoins := make(map[string]int)
  for i := 0; i < len(entity.players); i++ {
    playerCoins[entity.players[i].Name()] = entity.players[i].Coins()
  }
  return playerCoins
}

func (entity *Entity) AddCoins(number int) {
  entity.coins = entity.coins + number;
}
