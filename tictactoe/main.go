package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------------TICTACTOE---------------------------")
	fmt.Println("Example of a 3 x 3 game")
	thirdCol := [3]string{"2,2", "2,1", "2,0"}
	secondCol := [3]string{"1,2", "1,1", "1,0"}
	firstCol := [3]string{"0,2", "0,1", "0,0"}
	fmt.Println("\t ___________________________________")
	for i := 0; i < 3; i++ {
		fmt.Println("\t|           |           |           |")
		fmt.Println("\t|   ", firstCol[i], "   |   ", secondCol[i], "   |   ", thirdCol[i], "   |")
		fmt.Println("\t|___________|___________|___________|")
	}
	fmt.Println("")
	p1 := NewPlayer()
	for {
		fmt.Print("Enter name of fisrt Player : ")
		p1Name, _, err := reader.ReadLine()
		if err != nil {
			continue
		}

		if err := p1.SetName(string(p1Name)); err != nil {
			fmt.Println(err)
			continue
		}

		p1.SetMark("O")

		break
	}

	p2 := NewPlayer()
	for {
		fmt.Print("Enter name of second Player : ")
		p2Name, _, err := reader.ReadLine()
		if err != nil {
			continue
		}

		if err := p2.SetName(string(p2Name)); err != nil {
			fmt.Println(err)
			continue
		}

		p2.SetMark("X")

		break
	}

	var N int
	for {
		fmt.Print("N x N TicTacToe \nEnter number of N : ")
		nStr, _, err := reader.ReadLine()
		if err != nil {
			continue
		}

		nInt, err := strconv.Atoi(string(nStr))
		if err != nil {
			continue
		}

		if nInt <= 1 {
			fmt.Println("N is larger than 1")
			continue
		}

		N = nInt
		break
	}

	board := NewBoard(N)
	game := NewGame(p1, p2, board)
	fmt.Print("\n\n-----------------------Players Info-----------------------\n")
	fmt.Printf("\tPlayer 1 :- Name : %10s\tMark : %s\n", p1.Name, p1.Mark)
	fmt.Printf("\tPlayer 2 :- Name : %10s\tMark : %s\n", p2.Name, p2.Mark)
	fmt.Print("----------------------------------------------------------\n\n")
	fmt.Println()
	fmt.Println("-----------------------Game Starts------------------------")
	game.Play()
}
