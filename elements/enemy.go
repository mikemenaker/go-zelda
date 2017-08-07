package elements

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"go-zelda/utils"
	"math/rand"
	"time"
)

type Enemy struct {
	loc       pixel.Vec
	size      pixel.Rect
	direction int

	sheet       pixel.Picture
	anims       map[string][]pixel.Rect
	dying_sheet pixel.Picture
	dying_anims map[string][]pixel.Rect
	currFrame   pixel.Rect
	frameCount  int
	tick        int
	animCount   int

	isDying bool
	isDead  bool
}

const (
	DOWN = iota + 1
	UP
	LEFT
	RIGHT
)

func NewEnemy(loc pixel.Vec, name string) *Enemy {
	enemy := new(Enemy)

	sheet, anims, err := utils.LoadAnimationSheet("images/sprites/enemies/"+name+"/"+name+".png",
		"images/sprites/enemies/"+name+"/"+name+"_sheet.csv")
	if err != nil {
		panic(err)
	}

	dying_sheet, dying_anims, err := utils.LoadAnimationSheet("images/sprites/enemies/dying/dying.png", "images/sprites/enemies/dying/dying_sheet.csv")
	if err != nil {
		panic(err)
	}

	enemy.sheet = sheet
	enemy.anims = anims
	enemy.dying_sheet = dying_sheet
	enemy.dying_anims = dying_anims

	enemy.frameCount = 0
	enemy.tick = 0
	enemy.animCount = 0

	for anim := range anims {
		if anims[anim][0].Min.Y == 0 && anims[anim][0].Min.X == 0 {
			enemy.size = anims[anim][0]
			break
		}
	}

	enemy.loc = loc
	enemy.direction = DOWN

	return enemy
}

func (enemy *Enemy) setCurrentFrame() {
	if !enemy.isDead {
		frameKey := enemy.getFrameKey(enemy.direction)

		if enemy.isDying {
			if enemy.tick == 7 {
				enemy.tick = 0
				if enemy.frameCount == len(enemy.dying_anims[frameKey])-1 {
					enemy.frameCount = 0
					enemy.isDead = true
				} else {
					enemy.frameCount++
				}
			} else {
				enemy.tick++
			}
		} else {
			if enemy.tick == 7 {
				enemy.tick = 0
				if enemy.frameCount == len(enemy.anims[frameKey])-1 {
					enemy.frameCount = 0
					if enemy.animCount == 3 {
						enemy.updateDirection()
					} else {
						enemy.animCount++
					}
				} else {
					enemy.frameCount++
				}
			} else {
				enemy.tick++
			}
		}

		if enemy.isDying {
			enemy.currFrame = enemy.dying_anims[frameKey][enemy.frameCount]
		} else {
			enemy.currFrame = enemy.anims[frameKey][enemy.frameCount]
		}
	}
}

func (enemy *Enemy) getFrameKey(frameType int) string {
	if enemy.isDying {
		return "dying"
	}

	switch frameType {
	case LEFT:
		return "left"
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	case UP:
		return "up"
	}

	return "up"
}

func (enemy *Enemy) update(win *pixelgl.Window, objects []*Object, enemies []*Enemy) {
	if !enemy.isDying && !enemy.isDead {
		newPos := enemy.loc

		switch enemy.direction {
		case DOWN:
			newPos.Y -= .5
		case LEFT:
			newPos.X -= .5
		case UP:
			newPos.Y += .5
		case RIGHT:
			newPos.X += .5
		}

		overlapped := false
		bounds := utils.GetBounds(newPos, enemy.size)

		for _, o := range objects {
			if o.blocking && utils.Overlap(o.bounds, bounds) {
				overlapped = true
				break
			}
		}

		for _, e := range enemies {
			if enemy.loc != e.loc {
				enemyBounds := utils.GetBounds(e.loc, e.size)
				if utils.Overlap(enemyBounds, bounds) {
					overlapped = true
					break
				}
			}
		}

		screenBounds := pixel.R(enemy.size.Max.X*2, enemy.size.Max.Y*2, 1024-(enemy.size.Max.X*2), 768-(enemy.size.Max.Y*2))
		if !overlapped && utils.Overlap(bounds, screenBounds) {
			enemy.loc = newPos
			enemy.setCurrentFrame()
		} else {
			enemy.updateDirection()
		}
	} else if enemy.isDying {
		enemy.setCurrentFrame()
	}
}

func (enemy *Enemy) updateDirection() {
	enemy.frameCount = 0
	enemy.tick = 0
	enemy.animCount = 0

	direction := enemy.direction
	for direction == enemy.direction {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		direction = r1.Intn(RIGHT)
	}

	enemy.direction = direction
}

func (enemy *Enemy) draw(win *pixelgl.Window) {
	if !enemy.isDead {
		if enemy.isDying {
			sprite := pixel.NewSprite(enemy.dying_sheet, enemy.currFrame)
			sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(enemy.loc))
		} else {
			sprite := pixel.NewSprite(enemy.sheet, enemy.currFrame)
			sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(enemy.loc))
		}
	}
}
