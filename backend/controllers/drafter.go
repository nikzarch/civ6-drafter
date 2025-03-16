package controllers

import (
	"errors"
	"math/rand"
	"project/model"
)

type Draft struct {
	ID      int            `json:"id"`
	Nation  model.Nation   `json:"nation"`
	Leaders []model.Leader `json:"leaders"`
}

func GetDraft(players_amount int, choice_for_player int) ([][]Draft, error) {
	if players_amount*choice_for_player > len(model.Nations) {
		return nil, errors.New("Too much choice for too many players")
	}
	nations := make([]model.Nation, len(model.Nations))
	copy(nations, model.Nations[:])
	rand.Shuffle(len(nations), func(i, j int) {
		nations[i], nations[j] = nations[j], nations[i]
	})
	var result [][]Draft = make([][]Draft, players_amount)
	for i := 0; i < players_amount; i++ {
		for j := 0; j < choice_for_player; j++ {
			nation := nations[i*choice_for_player+j]
			result[i] = append(result[i], Draft{
				ID:      i*choice_for_player + j + 1,
				Nation:  nation,
				Leaders: nation.GetLeaders(),
			})
		}
	}
	return result, nil
}
