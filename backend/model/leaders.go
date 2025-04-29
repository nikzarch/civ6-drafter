package model

type Leader struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	DLC      string `json:"dlc"`
}
