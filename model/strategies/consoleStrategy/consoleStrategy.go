package consoleStrategy

import "github.com/twoodhouse/coup-sim/model/log"
import "github.com/twoodhouse/coup-sim/model/deck"
import "encoding/json"
import "fmt"

type Entity struct {
  strategyName string
  playerName string
}

func New() *Entity {
  var entity = Entity {
    "consoleStr",
    "notSet",
  }
  return &entity
}

func (entity *Entity) GetStrategyName() string {
  return entity.strategyName
}

func (entity *Entity) GetDuelFirstCardChoice() int {
  return 1
}

func (entity *Entity) SetPlayerName(name string) {
  entity.playerName = name
}

func (entity *Entity) GetAction(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) string {
  fmt.Println(log.PrettyJsonStr())
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var action string
  fmt.Printf("GetAction:\n> ")
  fmt.Scanf("%s\n", &action)
  return action
}

func (entity *Entity) GetLossChoice(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) int {
  fmt.Println(log.PrettyJsonStr())
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var choice int
  fmt.Printf("GetLossChoice:\n> ")
  fmt.Scanf("%d\n", &choice)
  return choice
}

func (entity *Entity) GetTarget(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) string {
  fmt.Println(log.PrettyJsonStr())
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var target string
  fmt.Printf("GetTarget:\n> ")
  fmt.Scanf("%s\n", &target)
  return target
}

func (entity *Entity) GetChallenge(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) bool {
  fmt.Println(log.PrettyJsonStr())
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var challenge bool
  fmt.Printf("GetChallenge:\n> ")
  fmt.Scanf("%t\n", &challenge)
  return challenge
}

func (entity *Entity) GetBlock(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) bool {
  fmt.Println(log.PrettyJsonStr())
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var block bool
  fmt.Printf("GetBlock:\n> ")
  fmt.Scanf("%t\n", &block)
  return block
}

func (entity *Entity) GetStealBlockCardChoice(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) int {
  fmt.Println(log.PrettyJsonStr())
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var choice int
  fmt.Printf("GetStealBlockCardChoice:\n> ")
  fmt.Scanf("%d\n", &choice)
  return choice
}

func (entity *Entity) GetExchangeKeepChoices(log *log.Entity, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) (int, int) {
  fmt.Println(log.PrettyJsonStr())
  // fmt.Printf("Cards gained by using exchange: %d%d\n", c1, c2)
  printPersonalTable(entity.playerName, playerNames, coinInfo, faceupInfo, deck)
  var r1 int
  fmt.Printf("GetExchangeKeepChoices(1):\n> ")
  fmt.Scanf("%d\n", &r1)
  var r2 int
  fmt.Printf("GetExchangeKeepChoices(2):\n> ")
  fmt.Scanf("%d\n", &r2)
  return r1, r2
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

func printPersonalTable(currentPlayerName string, playerNames []string, coinInfo map[string]int, faceupInfo map[string][]int, deck *deck.Entity) {
  fmt.Println("Player Name: " + currentPlayerName)
  fmt.Print("Your secret cards: ")
  for i := range deck.Cards() {
    fmt.Print(deck.Cards()[i])
  }
  fmt.Println()
  for i := range playerNames {
    fmt.Printf("%s: %d%d - %d coins\n", playerNames[i], faceupInfo[playerNames[i]][0], faceupInfo[playerNames[i]][1], coinInfo[playerNames[i]])
  }
}
