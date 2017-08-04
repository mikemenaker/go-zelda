package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"go-zelda/elements"
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

	intro := elements.NewIntro()
	link := elements.NewLink()
	world := elements.CreateWorld(elements.OVERWORLD)

	for !win.Closed() {
		if intro.IsActive {
			win.Clear(color.Black)
			intro.Update(win)
			intro.Draw(win)
		} else {
			world.UpdateAndDraw(win)
			worldType := link.Update(win, world)

			if worldType != elements.CURRENT {
				world = elements.CreateWorld(worldType)
				if world.LinkPos != pixel.V(0, 0) {
					link.Pos = world.LinkPos
				}
			} else {
				link.Draw(win)
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
