package main

import (
	"log"
	"os"

	"snake"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snakeParts := []snake.SnakePart{
		{
			X: 5,
			Y: 10,
		},
		{
			X: 6,
			Y: 10,
		},
		{
			X: 7,
			Y: 10,
		},
	}
	snakeBody := snake.SnakeBody{
		Parts:  snakeParts,
		Xspeed: 1,
		Yspeed: 0,
	}

	game := snake.Game{
		Screen:    screen,
		SnakeBody: snakeBody,
	}
	go game.Run()
	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.SnakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				game.SnakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				game.SnakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight {
				game.SnakeBody.ChangeDir(0, 1)
			}
		}
	}
}
