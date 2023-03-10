package main

import "errors"

type Player struct {
	Name string
	Mark string
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) SetName(name string) error {
	if len(name) > 10 {
		return errors.New("under 10 length")
	}
	p.Name = name

	return nil
}

func (p *Player) SetMark(mark string) error {
	if mark != "O" && mark != "X" {
		return errors.New("mark only O or X")
	}
	p.Mark = mark

	return nil
}
