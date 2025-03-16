package handlers

import (
	"encoding/json"
	"net/http"
	"project/controllers"
	"strconv"
)

func HandleDrafts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	players_amount, err_players := strconv.Atoi(r.URL.Query().Get("players"))
	choice_for_player, err_choices := strconv.Atoi(r.URL.Query().Get("choices"))
	if err_players != nil || err_choices != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	draft, err := controllers.GetDraft(players_amount, choice_for_player)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(draft)
}
