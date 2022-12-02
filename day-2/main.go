package main

import (
	"bytes"
	"log"
	"os"
)

type Moves struct {
	moveType    string
	movePoints  int
	losesToo    []string
	mapPointOne string
	mapPointTwo string
}

func (m Moves) IsALoss(move Moves) bool {
	lost := false
	for _, theMove := range m.losesToo {
		if theMove == move.moveType {
			lost = true
			break
		}
	}
	return lost
}

func (m Moves) IsSame(move Moves) bool {
	return m.moveType == move.moveType
}

func (m Moves) IsMove(playCode string) bool {
	return m.mapPointOne == playCode || m.mapPointTwo == playCode || false
}

type MoveCollection []Moves

func main() {
	input, err := os.ReadFile("input.txt")
	checkErr(err)

	splits := bytes.Split(input, []byte("\n"))

	rock := Moves{
		moveType:    "rock",
		movePoints:  1,
		losesToo:    []string{"paper"},
		mapPointOne: "A",
		mapPointTwo: "X",
	}

	paper := Moves{
		moveType:    "paper",
		movePoints:  2,
		losesToo:    []string{"scissors"},
		mapPointOne: "B",
		mapPointTwo: "Y",
	}

	scissors := Moves{
		moveType:    "scissors",
		movePoints:  3,
		losesToo:    []string{"rock"},
		mapPointOne: "C",
		mapPointTwo: "Z",
	}

	allMoves := MoveCollection{
		rock,
		paper,
		scissors,
	}

	score := 0

	for _, play := range splits {
		play = bytes.TrimSpace(play)

		if len(play) == 0 {
			continue
		}

		bothPlay := bytes.Split(play, []byte(" "))

		oppPlayCode := string(bothPlay[0])
		myPlayCode := string(bothPlay[1])

		var yourMove Moves
		var myMove Moves

		for _, guessedMove := range allMoves {
			if len(yourMove.moveType) == 0 && guessedMove.IsMove(oppPlayCode) {
				yourMove = guessedMove
			}
			if len(myMove.moveType) == 0 && guessedMove.IsMove(myPlayCode) {
				myMove = guessedMove
			}
		}

		if myMove.IsALoss(yourMove) {
			score += 0 + myMove.movePoints
		} else if myMove.IsSame(yourMove) {
			score += 3 + myMove.movePoints
		} else {
			score += 6 + myMove.movePoints
		}
	}

	log.Println(score)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
