package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const screenHeight = 240
const screenWidth = 320
var screenColor *color.Color

var players [4]*Player

func update(screen *ebiten.Image) error {

	//Fill the screen with #FF0000 color
	if screenColor == nil {
		randColor := randomColor()
		screenColor = &randColor
	}

	screen.Fill(*screenColor)

	// Update
	for _, p := range players {
		if p != nil {
			p.Update()
		}
	}

	// Draw
	for _, p := range players {
		if p != nil {
			p.Draw(screen)
		}
	}

	return nil
}

func main() {
	players[0] = NewPlayer(screenWidth/4, screenHeight/2, "P1")

	// Initialize Ebiten, and loop the update() function
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!"); err != nil {
		panic(err)
	}
}
