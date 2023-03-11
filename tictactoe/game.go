package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	PlayerOne *Player
	PlayerTwo *Player
	Board     *Board
	Turn      int
	Result    string
}

func NewGame(oneP, twoP *Player, board *Board) *Game {
	return &Game{
		PlayerOne: oneP,
		PlayerTwo: twoP,
		Board:     board,
		Turn:      board.N * board.N,
		Result:    "DRAW",
	}
}

func (g *Game) Play() {
	turn := g.Turn
	for turn > 0 {

		nowPlayer := g.PlayerOne
		if turn%2 == 0 {
			nowPlayer = g.PlayerTwo
		}

		X, Y, err := g.Input(nowPlayer.Name)
		if err != nil {
			fmt.Println("not integer")
			continue
		}
		if X > g.Board.N-1 || Y > g.Board.N-1 {
			fmt.Println("out of range")
			continue
		}
		if g.Board.isAlreadyMarked(X, Y) {
			fmt.Println("already marked")
			continue
		}

		g.Board.Mark(X, Y, nowPlayer.Mark)
		g.printBoard()
		isFinish := g.Board.CheckLineMadeUpSameChar(X, Y)
		turn--
		if isFinish {
			g.Result = nowPlayer.Name + " is Won"
			break
		}
	}

	g.printResult()
}

func (g *Game) Input(pName string) (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s Enter position : ", pName)
	bt, _, _ := reader.ReadLine()
	str := string(bt)
	strings.TrimSpace(str)
	strs := strings.Split(str, ",")
	X, err := strconv.Atoi(strs[0])
	if err != nil {
		return 0, 0, errors.New("is not Integer")
	}
	Y, err := strconv.Atoi(strs[1])
	if err != nil {
		return 0, 0, errors.New("is not Integer")
	}

	return X, Y, nil
}

func (g *Game) printBoard() {
	N := g.Board.N

	for i := 0; i < N; i++ {
		if i == 0 {
			fmt.Printf("\t")
			fmt.Print(" ________")
		} else if i == N-1 {
			fmt.Print("_______ ")
			fmt.Println("")
		} else {
			fmt.Print("________")
		}
	}
	for j := N - 1; j >= 0; j-- {
		for k := 0; k < 3; k++ {
			if k == 0 {
				for i := 0; i < N; i++ {
					if i == 0 {
						fmt.Print("\t")
					}
					fmt.Printf("|       ")

					if i == N-1 {
						fmt.Println("|")
					}
				}
			} else if k == 1 {
				for i := 0; i < N; i++ {
					if i == 0 {
						fmt.Printf("\t")
					}
					if g.Board.Board[i][j] == "" {
						fmt.Print("|       ")
					} else {
						fmt.Printf("|   %s   ", g.Board.Board[i][j])
					}

					if i == N-1 {
						fmt.Println("|")
					}
				}
			} else if k == 2 {
				for i := 0; i < N; i++ {
					if i == 0 {
						fmt.Printf("\t")
						fmt.Print("|_______|")
					} else if i == N-1 {
						fmt.Println("_______|")
					} else {
						fmt.Print("_______|")
					}
				}
			}
		}
	}
	fmt.Println("")
}

func (g *Game) printResult() {
	fmt.Println(g.Result)
}
