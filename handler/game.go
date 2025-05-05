package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PregOfficial/word-game/game"
	"github.com/PregOfficial/word-game/net"
	"github.com/google/uuid"
)

type StartResponse struct {
	ID        string    `json:"id"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type GuessResponse struct {
	Board  [][]game.LetterResult `json:"board"`
	IsOver bool                  `json:"isOver"`
	Won    bool                  `json:"won"`
}

var (
	games = make(map[string]*game.Game)
	mu    sync.Mutex
)

const maxGames = 1000

func StartGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ip := r.RemoteAddr
	if net.IsRateLimited(ip) {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	words, err := game.LoadWords("wordlist.txt")
	if err != nil {
		http.Error(w, "Server error, try again later", http.StatusServiceUnavailable)
		return
	}

	mu.Lock()
	if len(games) >= maxGames {
		mu.Unlock()
		http.Error(w, "Server busy, try again later", http.StatusServiceUnavailable)
		return
	}

	id := uuid.NewString()
	game := game.NewGame(id, words)
	games[game.Id] = &game
	mu.Unlock()

	time.AfterFunc(5*time.Minute, func() {
		mu.Lock()
		defer mu.Unlock()
		if _, exists := games[id]; exists {
			log.Printf("Auto-deleting game %s after 5 minutes", id)
			delete(games, id)
		}
	})

	log.Printf("Game with id %s started - secret %s", id, game.SecretWord)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(StartResponse{
		ID:        id,
		ExpiresAt: game.ExpiresAt,
	})
}

func Guess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID    string `json:"id"`
		Guess string `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON request", http.StatusBadRequest)
		return
	}

	req.Guess = strings.ToLower(strings.TrimSpace(req.Guess))

	if len(req.Guess) != 5 || !regexp.MustCompile(`^[a-z]{5}$`).MatchString(req.Guess) {
		http.Error(w, "guess must be exactly 5 lowercase letters (a-z)", http.StatusBadRequest)
		return
	}

	mu.Lock()
	game, ok := games[req.ID]
	mu.Unlock()

	if !ok {
		http.Error(w, "game not found", http.StatusNotFound)
		return
	}

	if game.IsOver {
		mu.Lock()
		delete(games, req.ID)
		mu.Unlock()
		http.Error(w, "game is over", http.StatusGone)
		return
	}

	game.Guess(req.Guess)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(GuessResponse{
		Board:  game.Guesses,
		IsOver: game.IsOver,
		Won:    game.Won,
	}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
