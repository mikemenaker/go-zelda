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

	pic, err := loadPicture("images/tree.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)
		sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(win.Bounds().Center()))
		//Vec, Rect
		link.update(win, pic.Bounds(), win.Bounds().Center())
		link.draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
