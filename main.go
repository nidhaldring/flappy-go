package main

import (
	"log"

	"github.com/gdamore/tcell"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := s.Init(); err != nil {
		log.Fatal(err.Error())
		return
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	loop := true
	for loop {
		s.Clear()
		s.SetContent(20, 22, 'A', nil, defStyle)
		s.SetContent(20, 25, 'B', nil, defStyle)
		s.SetContent(20, 26, 'C', nil, defStyle)
		s.SetContent(20, 28, 'D', nil, defStyle)
		s.Show()

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
				s.Fini()
				loop = false
			}
		}
	}

}
