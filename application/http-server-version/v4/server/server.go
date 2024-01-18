package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	PostPlayerScore(player string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
