package elements

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"go-zelda/utils"
)

type Door struct {
	Object
	target int
}

func NewDoor(img string, loc pixel.Vec, target int) *Door {
	door := new(Door)

	pic, err := utils.LoadPicture(img)
	if err != nil {
		panic(err)
	}

	door.size = pic.Bounds()
	door.loc = loc
	door.bounds = utils.GetBounds(door.loc, door.size)
	door.sprite = pixel.NewSprite(pic, door.size)
	door.blocking = false
	door.target = target

	return door
}

func (door *Door) Door(win *pixelgl.Window) {
	door.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(door.loc))
}
