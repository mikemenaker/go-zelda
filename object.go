package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Object struct {
	loc    pixel.Vec
	size   pixel.Rect
	bounds pixel.Rect
	sprite *pixel.Sprite
}

func NewObject(img string, loc pixel.Vec) *Object {
	object := new(Object)

	pic, err := loadPicture(img)
	if err != nil {
		panic(err)
	}

	object.size = pic.Bounds()
	object.loc = loc
	object.bounds = getBounds(object.loc, object.size)
	object.sprite = pixel.NewSprite(pic, object.size)

	return object
}

func (object *Object) draw(win *pixelgl.Window) {
	object.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(object.loc))
}
