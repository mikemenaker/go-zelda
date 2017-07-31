package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Door struct {
	Object
	target int
}

func NewDoor(img string, loc pixel.Vec, target int) *Door {
	door := new(Door)

	pic, err := loadPicture(img)
	if err != nil {
		panic(err)
	}

	door.size = pic.Bounds()
	door.loc = loc
	door.bounds = getBounds(door.loc, door.size)
	door.sprite = pixel.NewSprite(pic, door.size)
	door.blocking = false
	door.target = target

	return door
}

func (door *Door) Door(win *pixelgl.Window) {
	door.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(door.loc))
}
