package log

import "encoding/json"
import "strconv"
// import "fmt"

type Entity struct {
  jsonStr string
}

/*
We want data to be added dynamically to the structure
For example, a function CreateTurn should create one element in the array
*/

func New() *Entity {
  var entity = Entity {
    "[{\"playerName\": null, \"action\": {\"name\": null}}]",
  }
  return &entity
}

func (entity *Entity) JsonStr() string {
  return entity.jsonStr
}

func (entity *Entity) NextTurn() {
  entity.jsonStr = entity.jsonStr[:len(entity.jsonStr) - 1] + ",{\"playerName\": null, \"action\": {\"name\": null}}]"
}

func (entity *Entity) SetPlayerName(playerName string) {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  logTurn["playerName"] = playerName
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
}

func (entity *Entity) SetActionName(actionName string) {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  action := logTurn["action"].(map[string]interface{})
  action["name"] = actionName
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
}

func (entity *Entity) CreateTarget(target string) {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  action := logTurn["action"].(map[string]interface{})
  action["target"] = target
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
}

func (entity *Entity) CreateDisqualify(reason string) {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  logTurn["disqualified"] = reason
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
}

func (entity *Entity) CreateChallenge(challenger string, success bool, flippedCard int) {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  var successStr string
  if success {
    successStr = "true"
  } else {
    successStr = "false"
  }
  logTurn["challenge"] = unmarshalJson("{\"challenger\": \"" + challenger + "\", \"success\": " + successStr  + ", \"cardLoss\": " + strconv.Itoa(flippedCard) + "}")
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
}

func (entity *Entity) CreateBlock() {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  logTurn["block"] = unmarshalJson("{}")
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
}

func (entity *Entity) CreateBlockChallenge(challengeSuccess bool, flippedCard int) {
  log := unmarshalJsonArray(entity.jsonStr)
  logTurn := log[len(log) - 1]
  block := logTurn["block"].(map[string]interface{})
  var successStr string
  if challengeSuccess {
    successStr = "true"
  } else {
    successStr = "false"
  }
  block["challengeSuccess"] = successStr
  block["cardLoss"] = strconv.Itoa(flippedCard)
  marshaledJson, _ := json.Marshal(log)
  entity.jsonStr = (string(marshaledJson))
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

func (entity *Entity) PrettyJsonStr() string {
  log := unmarshalJsonArray(entity.jsonStr)
  prettyMarshaledJson, _ := json.MarshalIndent(log, "", "    ")
  return string(prettyMarshaledJson)
}
