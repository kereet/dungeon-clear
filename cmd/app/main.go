package main

import (
	"bufio"
	"dungeon-clear/internal/models"
	"dungeon-clear/internal/parser"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("data/events")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	game := models.NewGame()
	configData, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configData, &game)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		event, err := parser.ParseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(event.PlayerID, event.ID)
		player := game.Players[event.PlayerID]
		if player == nil {
			player = models.NewPlayer(event.PlayerID)
			game.Players[event.PlayerID] = player
			game.PlayersCount++
			if event.ID != 1 {
				player.IsDisqualified = true
				player.EntryDungeonTime = event.Time
				player.ExitDungeonTime = event.Time
				fmt.Printf("Player [%d] is disqualified\n", event.PlayerID)
			}
		}
		if player.IsDisqualified || player.IsDead {
			continue
		}
		switch event.ID {
		case 1:
			fmt.Printf("Player [%d] registered\n", player.ID)
		case 2:
			player.EntryDungeonTime = event.Time
			player.FloorEntryTime[1] = event.Time
			player.InDungeon = true
			fmt.Printf("Player [%d] entered the dungeon\n", player.ID)
		case 3:
			player.KilledMonsters[player.Floor]++
			fmt.Printf("Player [%d] killed the monster\n", player.ID)
			if player.KilledMonsters[player.Floor] == game.Floors {
				player.FloorClearTime[player.Floor] = event.Time
			}
		case 4:
			if player.Floor+1 <= game.Floors {
				player.Floor++
				player.FloorEntryTime[player.Floor] = event.Time
				fmt.Printf("Player [%d] went to the next floor\n", player.ID)
			} else {
				fmt.Printf("Player [%d] makes impossible move [4]\n", player.ID)
			}
		case 5:
			if player.Floor > 1 {
				player.Floor--
				fmt.Printf("Player [%d] went to the previous floor\n", player.ID)
			} else {
				fmt.Printf("Player [%d] makes impossible move [5]\n", player.ID)
			}
		case 6:
			player.Floor++
			player.BossEntryTime = event.Time
			fmt.Printf("Player [%d] entered the boss's floor\n", player.ID)
		case 7:
			fmt.Printf("Player [%d] killed the boss\n", player.ID)
			player.BossKilledTime = event.Time
		case 8:
			player.InDungeon = false
			player.ExitDungeonTime = event.Time
			fmt.Printf("Player [%d] left the dungeon\n", player.ID)
		case 9:
			fmt.Printf("Player [%d] cannot continue due to [%s]\n", player.ID, event.Extra)
			player.IsDisqualified = true
			player.ExitDungeonTime = event.Time
		case 10:
			health, err := strconv.Atoi(event.Extra)
			if err != nil {
				log.Fatal(err)
			}
			player.Health += health
			if player.Health > 100 {
				player.Health = 100
			}
			fmt.Printf("Player [%d] has restored [%s] of health\n", player.ID, event.Extra)
		case 11:
			damage, err := strconv.Atoi(event.Extra)
			if err != nil {
				log.Fatal(err)
			}
			player.Health -= damage
			fmt.Printf("Player [%d] received [%s] of damage\n", player.ID, event.Extra)
			if player.Health < 0 {
				fmt.Printf("Player [%d] is dead\n", player.ID)
				player.IsDead = true
				player.ExitDungeonTime = event.Time
				player.Health = 0
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Final report")
	for id := 1; id <= game.PlayersCount; id++ {
		player := game.Players[id]
		var status string
		avgDuration := "00:00:00"
		checkMonsters := true
		for _, monsters := range player.KilledMonsters {
			if monsters != game.Monsters {
				checkMonsters = false
				break
			}
		}
		if checkMonsters {
			status = "SUCCESS"
		} else {
			status = "FAIL"
		}
		if player.IsDisqualified {
			status = "DISQUAL"
		}
		if player.IsDead {
			status = "FAIL"
		}
		if checkMonsters {
			var sumDuration time.Duration
			for i := 1; i < game.Floors; i++ {
				sumDuration += player.FloorClearTime[i].Sub(player.FloorEntryTime[i])
			}
			avgDuration = formatDuration(sumDuration / time.Duration(game.Floors-1))
		}
		bossKillTime := "00:00:00"
		if !player.BossKilledTime.IsZero() {
			bossKillTime = formatDuration(player.BossKilledTime.Sub(player.BossEntryTime))
		}
		dungeonTime := formatDuration(player.ExitDungeonTime.Sub(player.EntryDungeonTime))
		fmt.Printf("[%s] %d [%s %s %s] HP: %d\n", status, id, dungeonTime, avgDuration, bossKillTime, player.Health)
	}
}

func formatDuration(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
