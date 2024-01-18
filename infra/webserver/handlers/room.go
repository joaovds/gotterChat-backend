package handlers

import (
	"encoding/json"
	"net/http"
)

type roomHandler struct {}

func NewRoomHandler() *roomHandler {
	return &roomHandler{}
}

type HubJson struct {
  Id string `json:"id"`
}

func (rh *roomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
  var HubsJson []HubJson
  for k, _ := range Hubs {
    HubsJson = append(HubsJson, HubJson{Id: k})
  }

  hubs, err := json.Marshal(HubsJson)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write(hubs)
}
