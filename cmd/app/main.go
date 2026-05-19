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
			game.Players[event.PlayerID].InDungeon = true
			fmt.Printf("Player [%d] entered the dungeon\n", player.ID)
		case 3:
			game.Players[player.ID].KilledMonsters[player.Floor]++
			fmt.Printf("Player [%d] killed the monster\n", player.ID)
		case 4:
			if player.Floor+1 <= game.Floors {
				player.Floor++
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
			fmt.Printf("Player [%d] entered the boss's floor\n", player.ID)
		case 7:
			fmt.Printf("Player [%d] killed the boss\n", player.ID)
		case 8:
			player.InDungeon = false
			fmt.Printf("Player [%d] left the dungeon\n", player.ID)
		case 9:
			fmt.Printf("Player [%d] cannot continue due to [%s]\n", player.ID, event.Extra)
			player.IsDisqualified = true
		case 10:
			health, err := strconv.Atoi(event.Extra)
			if err != nil {
				log.Fatal(err)
			}
			player.Health += health
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
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}
