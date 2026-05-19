package models

import "time"

type Game struct {
	Players      map[int]*Player
	PlayersCount int
	Floors       int       `json:"Floors"`
	Monsters     int       `json:"Monsters"`
	OpenAt       time.Time `json:"OpenAt"`
	Duration     time.Time `json:"Duration"`
}

func NewGame() *Game {
	return &Game{
		Players:      make(map[int]*Player),
		PlayersCount: 0,
	}
}
