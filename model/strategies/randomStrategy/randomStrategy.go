package randomStrategy

import "github.com/twoodhouse/coup-sim/model/log"
import "github.com/twoodhouse/coup-sim/model/deck"
import "encoding/json"
import "math/rand"
import "time"
// import "fmt"

type Entity struct {
  strategyName string
  playerName string
}

func New() *Entity {
  var entity = Entity {
    "RandomStr",
    "notSet",
  }
  rand.Seed(time.Now().UnixNano())
  return &entity
}

func (entity *Entity) GetStrategyName() string {
  return entity.strategyName
}

func (entity *Entity) GetDuelFirstCardChoice() int {
  return int(rand.Int31n(int32(5))) + 1
}

func (entity *Entity) SetPlayerName(name string) {
  entity.playerName = name
}

func (entity *Entity) GetAction(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) string {
  rand.Seed(time.Now().UnixNano())
  viableActions := make([]string, 0)
  viableActions = append(viableActions, "income")
  viableActions = append(viableActions, "foreign_aid")
  viableActions = append(viableActions, "tax")
  viableActions = append(viableActions, "steal")
  viableActions = append(viableActions, "exchange")
  if coinInfo[entity.playerName] >= 3 {
    viableActions = append(viableActions, "assassinate")
  }
  if coinInfo[entity.playerName] >= 7 {
    viableActions = append(viableActions, "coup")
  }

  if coinInfo[entity.playerName] >= 10 {
    return "coup"
  }
  return viableActions[rand.Int31n(int32(len(viableActions)))]
}

func (entity *Entity) GetLossChoice(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) int {
  if rand.Int31n(int32(2)) == 1 {
    return 0
  }
  return 1
}

func (entity *Entity) GetTarget(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) string {
  rand.Seed(time.Now().UnixNano())
  var choices []string
  for i := range playerNames {
    if playerNames[i] != entity.playerName && faceupInfo[playerNames[i]][1] == 0 {
      choices = append(choices, playerNames[i])
    }
  }
  return choices[rand.Int31n(int32(len(choices)))]
}

func (entity *Entity) GetChallenge(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) bool {
  rand.Seed(time.Now().UnixNano())
  if rand.Int31n(int32(2)) == 1 {
    return true
  }
  return false
}

func (entity *Entity) GetBlock(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) bool {
  rand.Seed(time.Now().UnixNano())
  if rand.Int31n(int32(2)) == 1 {
    return true
  }
  return false
}

func (entity *Entity) GetStealBlockCardChoice(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) int {
  if rand.Int31n(int32(2)) == 1 {
    return 2
  }
  return 5
}

func (entity *Entity) GetExchangeReturnChoices(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) (int, int) {
  choice1 := rand.Int31n(int32(deck.Size()))
  var choice2 int32
  for choice2 == choice1 {
    choice2 = rand.Int31n(int32(deck.Size()))
  }
  return deck.Cards()[choice1], deck.Cards()[choice2]
}

func unmarshalJsonArray(str string) []map[string]interface{} {
  byt := []byte(str)
  var dat []map[string]interface{}
  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  return dat
}

func unmarshalJson(str string) map[string]interface{} {
  byt := []byte(str)
  var dat map[string]interface{}
  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  return dat
}
