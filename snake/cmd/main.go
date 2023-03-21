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

	game := snake.Game{
		Screen: screen,
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
			} else if event.Key() == tcell.KeyUp && game.SnakeBody.Yspeed == 0 {
				game.SnakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyDown && game.SnakeBody.Yspeed == 0 {
				game.SnakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft && game.SnakeBody.Xspeed == 0 {
				game.SnakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight && game.SnakeBody.Xspeed == 0 {
				game.SnakeBody.ChangeDir(0, 1)
			} else if event.Rune() == 'y' && game.GameOver {
				go game.Run()
			} else if event.Rune() == 'n' && game.GameOver {
				game.Screen.Fini()
				os.Exit(0)
			}
		}
	}
}
