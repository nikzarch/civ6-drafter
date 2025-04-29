package services

import (
	"errors"
	"math/rand"
	"project/model"
)

type Draft struct {
	ID     int          `json:"id"`
	Nation model.Nation `json:"nation"`
	Leader model.Leader `json:"leader"`
}

func GetDraft(players_amount int, choice_for_player int) ([][]Draft, error) {
	nations, err := model.GetNations()
	if err != nil {
		return nil, err
	}
	if players_amount*choice_for_player > len(nations) {
		return nil, errors.New("Too much choice for too many players")
	}
	draft := make([]model.Nation, len(nations))
	copy(draft, nations[:])
	rand.Shuffle(len(draft), func(i, j int) {
		draft[i], draft[j] = draft[j], draft[i]
	})
	var result [][]Draft = make([][]Draft, players_amount)
	for i := 0; i < players_amount; i++ {
		for j := 0; j < choice_for_player; j++ {
			nation := draft[i*choice_for_player+j]
			result[i] = append(result[i], Draft{
				ID:     i*choice_for_player + j + 1,
				Nation: nation,
				Leader: nation.Leaders[rand.Intn(len(nation.Leaders))],
			})
		}
	}
	return result, nil
}
