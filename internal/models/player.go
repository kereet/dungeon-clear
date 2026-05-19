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
		TimeInDungeon: TimeInDungeon{
			FloorEntryTime: make(map[int]time.Time),
			FloorClearTime: make(map[int]time.Time),
		},
	}
}

type TimeInDungeon struct {
	EntryDungeonTime time.Time
	BossEntryTime    time.Time
	BossKilledTime   time.Time
	ExitDungeonTime  time.Time
	FloorEntryTime   map[int]time.Time
	FloorClearTime   map[int]time.Time
}
