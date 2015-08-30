package strategy
import "github.com/twoodhouse/coup-sim/model/log"

type Interface interface {
  GetName() string
  GetAction(log.Entity) string
  GetChallenge(log.Entity) bool
  GetBlock(log.Entity) bool
  GetAmbassadorReturns(log.Entity, string, string) (string,string)
}
