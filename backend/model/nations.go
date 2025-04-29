package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Nation struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Leaders  []Leader `json:"leaders"`
	ImageURL string   `json:"image_url"`
	DLC      string   `json:"dlc"`
}

var nations []Nation

func GetNations() ([]Nation, error) {
	if nations == nil {
		f, err := os.ReadFile("nations.json")
		if err != nil {
			return nil, errors.New("failed to read nations.json: " + err.Error())
		}
		if err := json.Unmarshal(f, &nations); err != nil {
			return nil, errors.New("failed to unmarshal nations.json: " + err.Error())
		}
	}
	fmt.Println(nations)
	return nations[:], nil
}
