package store

import "log"

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
