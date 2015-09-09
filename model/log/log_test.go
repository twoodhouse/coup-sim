package log
import (
  "testing"
  "fmt"
)

func TestLogCreation(t *testing.T) {
  log := New()
  log.SetPlayerName("Trevor")
  log.SetActionName("income")
  log.NextTurn()
  log.SetPlayerName("Michael")
  log.SetActionName("tax")
  log.CreateChallenge("Trevor", true)
  log.CreateChallengeCardLoss(5)
  log.NextTurn()
  log.SetPlayerName("Zeuterneuman")
  log.SetActionName("steal")
  log.CreateTarget("Trevor")
  log.NextTurn()
  log.SetPlayerName("Trevor")
  log.SetActionName("asssassinate")
  log.CreateTarget("Michael")
  log.CreateBlock()
  log.CreateBlockChallenge(true)
  log.CreateBlockChallengeCardLoss(1)
  log.CreateBlockCardClaim(2)
  log.CreateDisqualify("testing disqualify mechanism")
  log.CreateCardKilled(3)
  fmt.Println(log.PrettyJsonStr())
  // log.SetActionName("income")
  // log.NewTurn("Trevor")
}
