package main

import "fmt"

type Board struct {
	Board [][]string
	N     int
}

func NewBoard(N int) *Board {
	board := &Board{}
	board.Board = make([][]string, N)
	board.N = N
	for i := 0; i < N; i++ {
		board2 := make([]string, N)
		board.Board[i] = board2
	}

	return board
}

func (b *Board) CheckLineMadeUpSameChar(X, Y int) bool {
	var result bool
	if b.horizonCheck(X) || b.verticalCheck(Y) || b.diagonalCheck(X, Y) {
		result = true
	}

	return result
}

func (b *Board) horizonCheck(X int) bool {
	std := b.Board[X][0]
	cnt := 1
	for i := 1; i < b.N; i++ {
		if std == b.Board[X][i] {
			cnt++
		}
	}

	if cnt != b.N {
		return false
	}
	return true
}

func (b *Board) verticalCheck(Y int) bool {
	std := b.Board[0][Y]
	cnt := 1
	for i := 1; i < b.N; i++ {
		if std == b.Board[i][Y] {
			cnt++
		}
	}

	if cnt != b.N {
		return false
	}
	return true
}

func (b *Board) diagonalCheck(x int, y int) bool {
	a := x + y
	var result bool
	if a == b.N-1 && x == y {
		if b.leftToRightDigonal() || b.rightToLeftDigonal() {
			result = true
		}
	} else if a == b.N-1 {
		result = b.rightToLeftDigonal()
	} else if x == y {
		result = b.leftToRightDigonal()
	}

	return result
}

func (b *Board) leftToRightDigonal() bool {
	fmt.Println(b.Board)
	std := b.Board[0][0]
	cnt := 1
	for i := 1; i < b.N; i++ {
		if b.Board[i][i] == std {
			cnt++
		}
	}

	if cnt != b.N {
		return false
	}

	return true
}

func (b *Board) rightToLeftDigonal() bool {
	fmt.Println(b.Board)
	j := b.N - 1
	std := b.Board[0][j]
	cnt := 1

	for i := 1; i < b.N; i++ {
		j--
		if std == b.Board[i][j] {
			cnt++
		}
	}

	if cnt != b.N {
		return false
	}

	return true
}

func (b *Board) isAlreadyMarked(X, Y int) bool {
	if b.Board[X][Y] == "" {
		return false
	}
	return true
}

func (b *Board) Mark(X, Y int, mark string) {
	b.Board[X][Y] = mark
}
