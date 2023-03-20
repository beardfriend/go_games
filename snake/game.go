package snake

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen    tcell.Screen
	SnakeBody SnakeBody
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()

	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		g.Screen.Clear()
		g.SnakeBody.Update(width, height)
		g.Screen.SetContent(g.SnakeBody.X, g.SnakeBody.Y, ' ', nil, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
