package main

import "testing"

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
