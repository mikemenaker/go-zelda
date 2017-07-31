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

	intro := NewIntro()
	link := NewLink()
	world := createWorld(OVERWORLD)

	for !win.Closed() {
		if intro.isActive {
			win.Clear(color.Black)
			intro.update(win)
			intro.draw(win)
		} else {
			world.UpdateAndDraw(win)
			worldType := link.update(win, world)

			if worldType != CURRENT {
				world = createWorld(worldType)
			} else {
				link.draw(win)
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
