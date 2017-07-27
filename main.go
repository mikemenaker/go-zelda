package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
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

	// create intro
	intro := NewIntro()

	// init for overworld
	link := NewLink()
	objects, enemies := createWorld()
	bgColor := color.RGBA{72, 152, 72, 1}

	for !win.Closed() {
		if intro.isActive {
			win.Clear(bgColor)
			intro.update(win)
			intro.draw(win)
		} else {
			win.Clear(bgColor)

			for _, o := range objects {
				o.draw(win)
			}

			for _, e := range enemies {
				if !e.isDead {
					e.update(win, objects, enemies)
					e.draw(win)
				}
			}

			link.update(win, objects, enemies)
			link.draw(win)
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
