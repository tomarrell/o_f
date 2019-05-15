package main

// Response is the raw response returned
// by the onefootball API
type Response struct {
	Status string `json: "status"`
	Data   struct {
		Team Team `json: "team"`
	} `json: "data"`
}

// Team is the team structure within that response
type Team struct {
	ID      int         `json: "id"`
	Name    string      `json: "name"`
	Players []APIPlayer `json: "players"`
}

// APIPlayer is the player structure to be decoded
// by the API
type APIPlayer struct {
	ID        string `json: "id"`
	Age       string `json: "age"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

// Player also contains the teams that this
// player is a part of
type Player struct {
	ID        string
	Age       string
	Firstname string
	Lastname  string
	Teams     []string
}

// AddTeam adds another team to the player's Teams slice
func (s *Player) AddTeam(name string) {
	s.Teams = append(s.Teams, name)
}

// ByName to sort players based on their full name
type ByName []Player

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	iFullname := a[i].Firstname + a[i].Lastname
	jFullname := a[j].Firstname + a[j].Lastname
	return iFullname < jFullname
}
