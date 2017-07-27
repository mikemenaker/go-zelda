package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Intro struct {
	sprite *pixel.Sprite
	isActive      bool
}

func NewIntro() *Intro {
	intro := new(Intro)

	pic, err := loadPicture("images/intro_bg.png")
	if err != nil {
		panic(err)
	}

	intro.sprite = pixel.NewSprite(pic, pic.Bounds())
	intro.isActive = true
	return intro
}

func (intro *Intro) draw(win *pixelgl.Window) {
	intro.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(pixel.V(520, 520)))
}

func (intro *Intro) update(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeySpace) || win.JustPressed(pixelgl.KeySpace) {
		intro.isActive = false
	}
}
