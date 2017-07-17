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

	objects := createWorld()

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)

		for _, o := range objects {
			o.draw(win)
		}

		link.update(win, objects)
		link.draw(win)

		win.Update()
	}
}
func createWorld() []*Object {
	var objects []*Object
	objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384)))
	objects = append(objects, NewObject("images/tree.png", pixel.V(600, 384)))
	return objects
}

func main() {
	pixelgl.Run(run)
}
