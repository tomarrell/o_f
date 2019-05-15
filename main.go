package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// REQUESTS defines the total number of requests to be made against the
	// Onefootball API. Careful, this will also configure how thoroughly the
	// API is searched for the requested number of teams below.
	REQUESTS = 100
	// TIMEOUT defines the total number of seconds to wait before
	// timing out the channel, and moving on with the number of players
	// gathered by the API.
	TIMEOUT = 5

	// API defines the unformatted URL to request team information
	// given the team ID
	API = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
)

var clubs = []string{
	"Germany",
	"England",
	"France",
	"Spain",
	"Manchester United",
	"Arsenal",
	"Chelsea",
	"Barcelona",
	"Real Madrid",
	"Bayern Munich",
}

func main() {
	counter := 0
	c := make(chan Response)

	// Timeout the channel if all the clubs are not
	// found within the configured number of REQUESTS
	go func() {
		time.Sleep(TIMEOUT * time.Second)
		close(c)
	}()

	// Make parallel requests up to the limit
	// configured by REQUESTS
	for i := 0; i < REQUESTS; i++ {
		if counter == len(clubs) {
			break
		}

		// Format number into the URL as per docs
		url := fmt.Sprintf(API, i+1)
		go fetchClub(c, url)
	}

	// Pull the matched teams off the channel
	// and insert the players into a map to
	// build their teams slice
	playerMap := make(map[string]Player)

	for i := 0; i < len(clubs); i++ {
		res, ok := <-c
		if !ok {
			break
		}

		insertPlayers(playerMap, res.Data.Team)
		counter++
	}

	// Convert players into sortable slice
	var playerSlice []Player

	for _, player := range playerMap {
		playerSlice = append(playerSlice, player)
	}

	printPlayers(playerSlice)

	if counter < len(clubs) {
		fmt.Println("Failed to find all the clubs, try increasing the number of requests made")
	}
}

// Fetch the club and check if it is one
// that we are looking for
func fetchClub(c chan Response, url string) {
	var res Response

	resp, err := http.Get(url)
	if err != nil {
		panic("Failed to get data from API")
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println(err)
	}

	if contains(clubs, res.Data.Team.Name) {
		c <- res
	}
}

// Insert player in map with list of teams
func insertPlayers(m map[string]Player, team Team) {
	for _, player := range team.Players {
		if _, ok := m[player.ID]; ok {
			player := m[player.ID]
			player.AddTeam(team.Name)
			m[player.ID] = player
		} else {
			m[player.ID] = Player{
				ID:        player.ID,
				Age:       player.Age,
				Firstname: player.Firstname,
				Lastname:  player.Lastname,
				Teams:     []string{team.Name},
			}
		}
	}
}
