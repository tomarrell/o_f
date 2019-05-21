package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertPlayers(t *testing.T) {
	m := make(map[string]Player)
	team := Team{
		ID:   1,
		Name: "Team United",
		Players: []APIPlayer{
			{"1", "24", "Bob", "Jones"},
			{"2", "30", "James", "Lee"},
		},
	}

	team2 := Team{
		ID:   1,
		Name: "Team United",
		Players: []APIPlayer{
			{"1", "24", "Bob", "Jones"},
		},
	}

	insertPlayers(m, team)
	insertPlayers(m, team2)

	numTeams := len(m["1"].Teams)
	expected := 2
	if numTeams != expected {
		t.Errorf("Not enough teams for player with multiple associations got: %d expected: %d", numTeams, expected)
	}
}

func TestFetchClub(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/teams/en/23.json" {
			t.Error("URL not valid")
		}
	}))

	defer s.Close()

	fetchClub(22, s.URL)
}
