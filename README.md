# Word Guessing Game built in Go

Go test project with a React frontend.

My first Go project.

The frontend was built using GenAI.

## Instructions

#### Build

```
make build
```

#### Word List

The word list needs to be in the same directory as the binary

```
chmod +x word-guessing-game
cp wordlist.txt your-destination
```

#### Run

```
./word-guessing-game
```

## Endpoints

#### Frontend

```
http://localhost:8080
```

#### Start game

```
http://localhost:8080/api/start
```

#### Guess

```
http://localhost:8080/api/guess
```
