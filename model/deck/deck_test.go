package deck
import (
  "testing"
)

func TestDeckCreation(t *testing.T) {
  deck := New([]int{1,2,1,2,1,2,3,4,3,4,3,4,5,5,5})
  if !(deck.cards[1] == 2) {
    t.Errorf("Specific card wrong: expected %q, got %q", 0, deck.cards[0])
  }
}

func TestRandomDeckCreation(t *testing.T) {
  var deck = NewRandomCenter()
  if !(len(deck.cards) == 15){
    t.Errorf("Number of cards wrong: expected %d, got %d", 15, len(deck.cards))
  }
  if (deck.cards[0] == 0 && deck.cards[1] == 0 && deck.cards[2] == 0) {
    t.Errorf("Cards seem to be unrandomized")
  }
}

func TestTakeCards(t *testing.T) {
  deck := NewRandomCenter()
  if !(len(deck.cards) == 15){
    t.Errorf("Number of cards wrong: expected %d, got %d", 15, len(deck.cards))
  }
  cardsTaken := deck.TakeCards(2)
  deck.TakeCards(1)
  if !(len(deck.cards) == 12){
    t.Errorf("Number of cards wrong: expected %d, got %d", 12, len(deck.cards))
  }
  if !(len(cardsTaken) == 2){
    t.Errorf("Number of cards taken wrong: expected %d, got %d", 3, len(cardsTaken))
  }
}

func TestGiveCards(t *testing.T) {
  deck := NewRandomCenter()
  if !(len(deck.cards) == 15){
    t.Errorf("Number of cards wrong: expected %d, got %d", 15, len(deck.cards))
  }
  deck.GiveCards([]int{2, 3})
  if !(len(deck.cards) == 17){
    t.Errorf("Number of cards wrong: expected %d, got %d", 17, len(deck.cards))
  }
}
