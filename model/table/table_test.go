package table
import (
	// "fmt"
	"testing"
	"github.com/twoodhouse/coup-sim/model/player"
	"github.com/twoodhouse/coup-sim/model/deck"
)

// func TestTableCreation(t *testing.T) {
// 	var playerDeck = deck.New("pdeck", false)
// 	var table = New("table1", []player.Entity{player.New("p1", playerDeck, 2), player.New("p2", playerDeck, 2)}, deck.New("d1", true), 5)
// 	if !(table.Name() == "table1") {
// 		t.Errorf("Table name wrong: expected %q, got %q", "table1", table.Name())
// 	}
// 	if !(table.Players()[0].Name() == "p1") {
// 		t.Errorf("Player name wrong: expected %q, got %q", "p1", table.Players()[0].Name())
// 	}
// 	if !(table.Players()[1].Name() == "p2") {
// 		t.Errorf("Player name wrong: expected %q, got %q", "p2", table.Players()[1].Name())
// 	}
// 	if !(table.Decks()[0].Name() == "d1") {
// 		t.Errorf("Deck name wrong: expected %q, got %q", "d1", table.Decks()[0].Name())
// 	}
// 	if !(table.Coins() == 5) {
// 		t.Errorf("Number of coins wrong: expected %q, got %q", 5, table.Coins())
// 	}
// }

func TestAddCoins(t *testing.T) {
	var table = New("table1", []player.Entity{player.New("p1", 2), player.New("p2", 2)}, deck.New("d1", true), 5)
	table.AddCoins(5)
	if !(table.Coins() == 10) {
		t.Errorf("Number of coins wrong: expected %d, got %d", 10, table.Coins())
	}
	table.AddCoins(-15)
	if !(table.Coins() == -5) {
		t.Errorf("Number of coins wrong: expected %d, got %d", -5, table.Coins())
	}
}
// fmt.Printf(table.name);
// []player.Entity{{"p1"}, {"p2"}}

// cases := []struct {
// 	want string
// }{
// 	{"Hello, world", "dlrow ,olleH},
// 	{"Hello, 世界", "界世 ,olleH"},
// 	{"", ""},
// }
// for _, c := range cases {
// 	got := Reverse(c.in)
// 	if got != c.want {
// 		t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
// 	}
// }
