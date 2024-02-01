package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	PostPlayerScore(player string)
}

type PlayerServer struct {
	Store   PlayerStore
	Handler http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.Store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 4},
	}
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case "GET":
		p.getPlayerScore(w, player)

	case "POST":
		p.postPlayerScore(w, player)
	}
}

func (p *PlayerServer) getPlayerScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) postPlayerScore(w http.ResponseWriter, player string) {
	p.Store.PostPlayerScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(player string) int {
	if player == "Artus" {
		return 20
	}

	if player == "Amanda" {
		return 10
	}

	return 0
}
