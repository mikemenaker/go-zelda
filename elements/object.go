package elements

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"go-zelda/utils"
)

type Object struct {
	loc      pixel.Vec
	size     pixel.Rect
	bounds   pixel.Rect
	sprite   *pixel.Sprite
	blocking bool
}

func NewObject(img string, loc pixel.Vec, blocking bool) *Object {
	object := new(Object)

	pic, err := utils.LoadPicture(img)
	if err != nil {
		panic(err)
	}

	object.size = pic.Bounds()
	object.loc = loc
	object.bounds = utils.GetBounds(object.loc, object.size)
	object.sprite = pixel.NewSprite(pic, object.size)
	object.blocking = blocking

	return object
}

func (object *Object) draw(win *pixelgl.Window) {
	object.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(object.loc))
}
