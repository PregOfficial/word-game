# Word Guessing Game built in Go

Go test project with a React frontend.

The frontend was built with the help of GenAI.

[![Build and Upload](https://github.com/PregOfficial/word-game/actions/workflows/build.yaml/badge.svg)](https://github.com/PregOfficial/word-game/actions/workflows/build.yaml)

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

Returns

```json
{
    "id": string,
    "expiresAt": Date
}
```

#### Guess

```sh
http://localhost:8080/api/guess

curl --location 'localhost:8080/api/guess' \
--header 'Content-Type: application/json' \
--data '{
    "id": "uuid",
    "guess": "hello"
}'
```

Returns

```json
{
    "board": [
        [
            {
                "letter": "h",
                "color": "yellow"
            },
            {
                "letter": "e",
                "color": "yellow"
            },
            {
                "letter": "l",
                "color": "gray"
            },
            {
                "letter": "l",
                "color": "gray"
            },
            {
                "letter": "0",
                "color": "green"
            }
        ],
        ...
    ],
    "isOver": false,
    "won": false
}
```
