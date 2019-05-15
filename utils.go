package main

import (
	"fmt"
	"sort"
	"strings"
)

// Check if a string value to check is contained
// within a string slice
func contains(s []string, check string) bool {
	for _, val := range s {
		if val == check {
			return true
		}
	}

	return false
}

// PrintPlayers outputs the players in the given format
func printPlayers(players []Player) {
	sort.Sort(ByName(players))

	for i, p := range players {
		fmt.Println(formatPlayer(i, p))
	}
}

func formatPlayer(i int, p Player) string {
	fullname := p.Firstname + p.Lastname
	return fmt.Sprintf("%d. %s; %s; %s", i+1, fullname, p.Age, strings.Join(p.Teams, ", "))
}
