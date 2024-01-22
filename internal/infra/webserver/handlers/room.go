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
  if len(Hubs) == 0 {
    w.Write([]byte("[]"))
    return
  }

  var HubsJson []HubJson
  for id := range Hubs {
    HubsJson = append(HubsJson, HubJson{Id: id})
  }

  hubs, err := json.Marshal(HubsJson)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write(hubs)
}
