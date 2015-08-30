package table
import (
	"testing"
	"github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
	"github.com/twoodhouse/coup-sim/model/strategy"
)

func TestTableCreation(t *testing.T) {
  s1 := noLieStrategy.New()
	s2 := noLieStrategy.New()
	s3 := noLieStrategy.New()
	strategies := make([]strategy.Interface, 3)
	strategies[0] = &s1
	strategies[1] = &s2
	strategies[2] = &s3
	table := New(strategies)

	if !(table.players[0].DeckSize() == 2) {
		t.Errorf("Player deck size wrong: expected %d, got %d", 2, table.players[0].DeckSize())
	}
	if !(table.players[2].DeckSize() == 2) {
		t.Errorf("Player deck size wrong: expected %d, got %d", 2, table.players[2].DeckSize())
	}
	if !(table.centerDeck.Size() == 9) {
		t.Errorf("Center deck size wrong: expected %d, got %d", 9, table.centerDeck.Size())
	}
}
