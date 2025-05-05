.PHONY: all build frontend backend clean

all: build

build: frontend backend

frontend:
	cd word-game-web && npm install && npm run build

backend: backend-mac backend-linux backend-windows

backend-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/word-guessing-game-mac

backend-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/word-guessing-game-linux

backend-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/word-guessing-game.exe

clean:
	rm -f bin/word-guessing-game* word-guessing-game
	cd word-game-web && rm -rf dist node_modules
