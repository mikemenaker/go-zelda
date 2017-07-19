package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Enemy struct {
	loc    pixel.Vec
	size   pixel.Rect
	bounds pixel.Rect
	sprite *pixel.Sprite
}

func NewEnemy(img string, loc pixel.Vec) *Enemy {
	enemy := new(Enemy)

	pic, err := loadPicture(img)
	if err != nil {
		panic(err)
	}

	enemy.size = pic.Bounds()
	enemy.loc = loc
	enemy.bounds = getBounds(enemy.loc, enemy.size)
	enemy.sprite = pixel.NewSprite(pic, enemy.size)

	return enemy
}

func (enemy *Enemy) draw(win *pixelgl.Window) {
	enemy.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(enemy.loc))
}
