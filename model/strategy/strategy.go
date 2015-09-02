package strategy
import "github.com/twoodhouse/coup-sim/model/log"
import "github.com/twoodhouse/coup-sim/model/deck"

type Interface interface {
  GetName() string
  GetAction(*log.Entity, map[string]int, map[string][]int, *deck.Entity) string
  GetLossChoice(*log.Entity, map[string]int, map[string][]int, *deck.Entity) int
  GetChallenge(*log.Entity, map[string]int, map[string][]int, *deck.Entity) bool
  GetBlock(*log.Entity, map[string]int, map[string][]int, *deck.Entity) bool
  GetAmbassadorReturns(*log.Entity, map[string]int, map[string][]int, string, string, *deck.Entity) (string,string)
}
