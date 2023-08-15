package main

import (
	"math"
	"time"

	"github.com/gdamore/tcell"
)

// reference center of the bird
var (
	// flyingBird  = "/\\"
	// fallingBird = "\\/"
	birdX     = 10.0
	birdY     = 10.0
	isJumping = false
)

func DrawBird(s tcell.Screen) {
	birdY += 0.0001
	if isJumping {
		birdX += 0.0002
		birdY = math.Sin(birdX) / 1000
	}

	s.SetContent(int(birdX), int(birdY), '/', nil, DefStyle)
	s.SetContent(int(birdX+1), int(birdY), '\\', nil, DefStyle)
}

func MakeBirdJump() {
	if isJumping {
		return
	}

	isJumping = true
	time.AfterFunc(time.Second, func() {
		isJumping = false
	})
}
