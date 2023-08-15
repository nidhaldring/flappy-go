package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

var DefStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

func InitScreen() (tcell.Screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := s.Init(); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	s.SetStyle(DefStyle)
	return s, nil
}

func OnEvents(s tcell.Screen, mainLoop chan<- bool) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
				mainLoop <- true
			} else if ev.Key() == tcell.KeyUp {
				MakeBirdJump()
			}
		}
	}
}

func DrawingLoop(s tcell.Screen, mainLoopHasEnded <-chan bool) {
	for {
		s.Clear()
		DrawBird(s)
		s.Show()
	}
}

func main() {
	s, err := InitScreen()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer s.Fini()

	mainLoopHasEnded := make(chan bool)

	go OnEvents(s, mainLoopHasEnded)
	go DrawingLoop(s, mainLoopHasEnded)

	<-mainLoopHasEnded
}
