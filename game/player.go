package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/gofont/goregular"
)

type Player struct {
	eImg *ebiten.Image
	X, Y int
	height, width int
	title string
	playerColor color.Color
}

func NewPlayer(x, y int, title string) *Player {
	height := 16
	width := 16

	// Create an image
	playerImg, _ := ebiten.NewImage(width, height, ebiten.FilterNearest)

	// Fill the square with the white color
	//playerImg.Fill(color.White)
	playerColor := randomColor()
	playerImg.Fill(playerColor)

	return &Player{
		title: title,
		X: x,
		Y: y,
		height: height,
		width: width,
		eImg: playerImg,
		playerColor: playerColor,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw the square image to the screen with an empty option
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.X- p.width /2), float64(p.Y - p.height /2))
	screen.DrawImage(p.eImg, opts)

	font, _ := truetype.Parse(goregular.TTF)

	text.Draw(screen, p.title, truetype.NewFace(font, nil), p.X-(p.width / 2), p.Y - p.height /2 - 5, p.playerColor)
}

func (p *Player) Update() {
	x, y := ebiten.CursorPosition()
	p.X = x
	p.Y = y

	{	// Don't exceed boundaries
		if p.X+p.width/2 > screenWidth {
			p.X = screenWidth - p.width/2
		}

		if p.X-p.width/2 < 0 {
			p.X = 0 + p.width/2
		}

		if p.Y+p.height/2 > screenHeight {
			p.Y = screenHeight - p.height/2
		}

		if p.Y-p.height/2 < 0 {
			p.Y = 0 + p.height/2
		}
	}
}

func randomColor() color.Color {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	r := uint8(r1.Intn(255))
	g := uint8(r1.Intn(255))
	b := uint8(r1.Intn(255))

	return color.RGBA{r, g, b, 255}
}