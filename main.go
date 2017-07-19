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

	link := NewLink()

	objects, enemies := createWorld()

	bgColor := color.RGBA{72, 152, 72, 1}
	for !win.Closed() {
		win.Clear(bgColor)

		for _, o := range objects {
			o.draw(win)
		}

		for _, e := range enemies {
			e.draw(win)
		}

		link.update(win, objects, enemies)
		link.draw(win)

		win.Update()
	}
}
func createWorld() ([]*Object, []*Enemy) {
	var objects []*Object
	objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384)))
	objects = append(objects, NewObject("images/tree.png", pixel.V(600, 384)))

	var enemies []*Enemy
	enemies = append(enemies, NewEnemy("images/green_soldier.png", pixel.V(440, 384)))

	return objects, enemies
}

func main() {
	pixelgl.Run(run)
}
