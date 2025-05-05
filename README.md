# Word Guessing Game built in Go

Go test project with a React frontend.

The frontend was built with the help of GenAI.

## Instructions

#### Build

```
make build
```

#### Word List

The word list needs to be in the same directory as the binary

```
chmod +x word-guessing-game-<<platform>>
cp wordlist.txt your-destination
```

#### Run

```
./word-guessing-game-<<platform>>
```

## Endpoints

#### Frontend

```sh
http://localhost:8080
```

#### Start game

```sh
http://localhost:8080/api/start

curl --location --request POST 'localhost:8080/api/start'
```

#### Guess

```sh
http://localhost:8080/api/guess
```
