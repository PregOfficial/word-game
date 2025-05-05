.PHONY: all build frontend backend clean

all: build

build: frontend backend

frontend:
	cd word-game-web && npm install && npm run build

backend:
	go build -o word-guessing-game

clean:
	rm -f word-guessing-game
	cd word-game-web && rm -rf dist node_modules
