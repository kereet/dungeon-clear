package parser

import (
	"dungeon-clear/internal/models"
	"strconv"
	"strings"
	"time"
)

func ParseLine(line string) (event *models.Event, err error) {
	arr := strings.Split(line, " ")
	timeStr, playerID, eventID := arr[0], arr[1], arr[2]
	var extra string
	if len(arr) == 4 {
		extra = arr[3]
	}
	t, err := time.Parse("15:04:05", strings.Trim(timeStr, "[]"))
	if err != nil {
		return nil, err
	}
	intEventID, err := strconv.Atoi(eventID)
	if err != nil {
		return nil, err
	}
	intPlayerID, err := strconv.Atoi(playerID)
	if err != nil {
		return nil, err
	}
	return models.NewEvent(intEventID, intPlayerID, t, extra), nil
}
