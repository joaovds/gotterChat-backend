package handlers

import (
	"encoding/json"
	"net/http"
)

type roomHandler struct{}

func NewRoomHandler() *roomHandler {
	return &roomHandler{}
}

type HubJson struct {
	Id string `json:"id"`
}

// Get All Rooms godoc
// @Summary Get All Rooms
// @Description Get All Rooms
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {array} HubJson
// @Failure 500 {string} string "Internal Server Error"
// @Router /rooms [get]
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
