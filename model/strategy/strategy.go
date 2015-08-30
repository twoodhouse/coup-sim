package strategy
import "github.com/twoodhouse/coup-sim/model/log"

type Interface interface {
  GetName() string
  GetAction(log.Entity, map[string]int, map[string][]int) string
  GetChallenge(log.Entity, map[string]int, map[string][]int) bool
  GetBlock(log.Entity, map[string]int, map[string][]int) bool
  GetAmbassadorReturns(log.Entity, map[string]int, map[string][]int, string, string) (string,string)
}
