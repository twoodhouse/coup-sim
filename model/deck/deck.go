package deck
// import "fmt"
import "math/rand"
import "time"

type Entity struct {
  cards []int
}

func New(cards []int) *Entity {
  var entity = Entity {
    cards,
  }
  return &entity
}

func NewRandomCenter() *Entity {
  var cards = []int{5,5,5,1,1,1,2,2,2,3,3,3,4,4,4}
  var entity = Entity {
    cards,
  }
  entity.ShuffleCards()
  return &entity
}

func (entity *Entity) TakeCards(number int) []int {
  cardsRemaining := make([]int, entity.Size() - number)
  for i := 0; i < len(entity.cards) - number; i++ {
    cardsRemaining[i] = entity.cards[i]
  }
  cardsTaken := make([]int, number)
  for i := 0; i < number; i++ {
    cardsTaken[i] = entity.cards[len(entity.cards) - i - 1]
  }
  entity.cards = cardsRemaining
  return cardsTaken
}

func (entity *Entity) TakeTopCard() int {
  cardsRemaining := make([]int, entity.Size() - 1)
  for i := 0; i < len(entity.cards) - 1; i++ {
    cardsRemaining[i] = entity.cards[i]
  }
  cardTaken := entity.cards[len(entity.cards) - 1]
  entity.cards = cardsRemaining
  return cardTaken
}

func (entity *Entity) TakeBottomCard() int {
  cardsRemaining := make([]int, entity.Size() - 1)
  for i := 0; i < len(entity.cards) - 1; i++ {
    cardsRemaining[i] = entity.cards[i + 1]
  }
  cardTaken := entity.cards[0]
  entity.cards = cardsRemaining
  return cardTaken
}

func (entity *Entity) GiveCards(givenCards []int) {
  newCards := make([]int, len(entity.cards) + len(givenCards))
  for i := 0; i < len(entity.cards); i++ {
    newCards[i] = entity.cards[i]
  }
  for i := 0; i < len(givenCards); i++ {
    newCards[i + len(entity.cards)] = givenCards[i]
  }
  entity.cards = newCards
  entity.ShuffleCards()
}

func (entity *Entity) ShuffleCards() {
  inbound := entity.cards
  dest := make([]int, len(inbound))
  rand.Seed(time.Now().Unix())
  perm := rand.Perm(len(inbound))
  for i, v := range perm {
      dest[v] = inbound[i]
  }
  entity.cards = dest
}

func (entity *Entity) Size() int {
  return len(entity.cards)
}

func (entity *Entity) Cards() []int {
  return entity.cards
}

func (entity *Entity) HasCardForAction(cardActionName string) bool {
  hasCard := false
  var cardValue int
  switch cardActionName {
  case "tax":
    cardValue = 1
  case "steal":
    cardValue = 2
  case "assassinate":
    cardValue = 3
  case "block":
    cardValue = 4
  case "exchange":
    cardValue = 5
  }
  for i := 0; i < len(entity.cards); i++ {
    if entity.cards[i] == cardValue {
      hasCard = true
    }
  }
  return hasCard
}

func (entity *Entity) HasCards(c1 int, c2 int) bool {
  c1Found := false
  c2Found := false
  for i := 0; i < len(entity.cards); i++ {
    if entity.cards[i] == c1 && !c1Found {
      c1Found = true
    } else if entity.cards[i] == c2 && !c2Found {
      c2Found = true
    }
  }
  return c1Found && c2Found
}

func (entity *Entity) ReplaceCardWithCard(oldCard int, newCard int) {
  if entity.Cards()[0] == oldCard {
    entity.Cards()[0] = newCard
  } else if entity.Cards()[1] == oldCard {
    entity.Cards()[1] = newCard
  }
}
