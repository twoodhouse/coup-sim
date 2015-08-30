package deck
// import "fmt"
import "math/rand"
import "time"

type Entity struct {
  cards []int
}

func New(cards []int) Entity {
  var entity = Entity {
    cards,
  }
  return entity
}

func NewRandomCenter() Entity {
  var cards = []int{0,0,0,1,1,1,2,2,2,3,3,3,4,4,4}
  var entity = Entity {
    cards,
  }
  entity.shuffleCards()
  return entity
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

func (entity *Entity) GiveCards(givenCards []int) {
  newCards := make([]int, len(entity.cards) + len(givenCards))
  for i := 0; i < len(entity.cards); i++ {
    newCards[i] = entity.cards[i]
  }
  for i := 0; i < len(givenCards); i++ {
    newCards[i + len(entity.cards)] = entity.cards[i]
  }
  entity.cards = newCards
  entity.shuffleCards()
}

func (entity *Entity) shuffleCards() {
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
