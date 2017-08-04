package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Intro struct {
	background *pixel.Sprite
	text       *pixel.Sprite
	isActive   bool
}

func NewIntro() *Intro {
	intro := new(Intro)

	pic, err := loadPicture("images/intro/intro_bg.png")
	if err != nil {
		panic(err)
	}
	intro.background = pixel.NewSprite(pic, pic.Bounds())

	pic, err = loadPicture("images/intro/intro_text.png")
	if err != nil {
		panic(err)
	}
	intro.text = pixel.NewSprite(pic, pic.Bounds())

	intro.isActive = true
	return intro
}

func (intro *Intro) draw(win *pixelgl.Window) {
	intro.background.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(pixel.V(520, 520)))
	intro.text.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(pixel.V(520, 520)))
}

func (intro *Intro) update(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeySpace) || win.JustPressed(pixelgl.KeySpace) {
		intro.isActive = false
	}
}
