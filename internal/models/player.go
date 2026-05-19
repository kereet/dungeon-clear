package models

import "time"

type Player struct {
	ID             int
	Status         string
	Health         int
	Floor          int
	KilledMonsters map[int]int
	InDungeon      bool
	TimeInDungeon
	IsDisqualified bool
	IsDead         bool
}

func NewPlayer(id int) *Player {
	return &Player{
		ID:             id,
		Health:         100,
		Floor:          1,
		KilledMonsters: make(map[int]int),
	}
}

type TimeInDungeon struct {
	EntryDungeonTime time.Time
	BossEntryTime    time.Time
	BossKilledTime   time.Time
	ExitDungeonTime  time.Time
	FloorEntry       []int
	FloorExit        []int
}
