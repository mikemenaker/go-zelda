package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Go Zelda",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	link := NewLink()

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)

		link.update(win)
		link.draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
