package table
import (
  "fmt"
  "github.com/twoodhouse/coup-sim/model/player"
  "github.com/twoodhouse/coup-sim/model/strategy"
  "github.com/twoodhouse/coup-sim/model/deck"
)

type Entity struct {
  players []*player.Entity
  centerDeck *deck.Entity
  faceupDecks []*deck.Entity
  coins int
}

func New(strategies []strategy.Interface, playerNames []string) *Entity {
  centerDeck := deck.NewRandomCenter()
  faceupDecks := make([]*deck.Entity, len(strategies))

  players := make([]*player.Entity, len(strategies))
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
  return &entity
}

func (entity *Entity) Players() []*player.Entity {
  return entity.players
}

func (entity *Entity) PlayerNames() []string {
  names := make([]string, len(entity.players))
  for i := range entity.players {
    names[i] = entity.players[i].Name()
  }
  return names
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

func (entity *Entity) CenterDeck() *deck.Entity {
  return entity.centerDeck
}

func (entity *Entity) AddCoins(number int) {
  entity.coins = entity.coins + number;
}

func (entity *Entity) PrintTable() {
  for i := 0; i < len(entity.players); i++ {
    card1 := entity.faceupDecks[i].Cards()[0]
    card2 := entity.faceupDecks[i].Cards()[1]
    fmt.Printf("%s: %d%d, %d coins\n", entity.players[i].Name(), card1, card2, entity.players[i].Coins())
  }
}
