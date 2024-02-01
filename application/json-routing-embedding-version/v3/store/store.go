package store

import (
	"log"

	srv "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v3/server"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	log.Printf("Got score of player %s\n", player)
	return i.store[player]
}

func (i *InMemoryPlayerStore) PostPlayerScore(player string) {
	i.store[player]++
	log.Printf("Posted score %v to player %s\n", i.store[player], player)
}

func (i *InMemoryPlayerStore) GetLeague() []srv.Player {
	var league []srv.Player

	for name, wins := range i.store {
		league = append(league, srv.Player{Name: name, Wins: wins})
	}
	log.Printf("Got player league %v\n", league)
	return league
}
