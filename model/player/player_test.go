package player
import (
  "testing"
  "github.com/twoodhouse/coup-sim/model/deck"
)

func TestPlayerCreation(t *testing.T) {
  var player = New("p1", deck.NewRandomCenter(), 2)
  if !(player.Name() == "p1") {
    t.Errorf("Table name wrong: expected %q, got %q", "p1", player.Name())
  }
}

func TestPlayerDeckCreation(t *testing.T) {
  centerDeck := deck.NewRandomCenter()
  playerDeck := deck.New(centerDeck.TakeCards(2))
  if !(playerDeck.Size() == 2) {
    t.Errorf("Player deck size wrong: expected %d, got %d", 2, playerDeck)
  }
  if !(centerDeck.Size() == 13) {
    t.Errorf("Player deck size wrong: expected %d, got %d", 13, centerDeck)
  }
}
