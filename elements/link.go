package elements

import (
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"go-zelda/utils"
)

type Link struct {
	sheet         pixel.Picture
	anims         map[string][]pixel.Rect
	Pos           pixel.Vec
	currFrame     pixel.Rect
	lastFrameType int
	frameCount    int
	tick          int
}

func NewLink() *Link {
	link := new(Link)

	sheet, anims, err := utils.LoadAnimationSheet("images/sprites/link/link.png", "images/sprites/link/sheet.csv")
	if err != nil {
		panic(err)
	}
	link.sheet = sheet
	link.anims = anims

	link.Pos = pixel.V(130, 130)
	link.currFrame = pixel.R(0, 0, 0, 0)

	link.frameCount = 0
	link.lastFrameType = STAND
	link.tick = 0
	return link
}

const (
	WALK_LEFT = iota + 1
	WALK_RIGHT
	WALK_DOWN
	WALK_UP
	STAND
	ATTACK_UP
	ATTACK_DOWN
	ATTACK_LEFT
	ATTACK_RIGHT
)

func (link *Link) setCurrentFrame(frameType int) {
	frameKey := link.getFrameKey(frameType)

	if frameType != link.lastFrameType {
		link.lastFrameType = frameType
		link.frameCount = 0
		link.tick = 0
	} else {
		if link.tick == 3 {
			link.tick = 0
			if link.frameCount == len(link.anims[frameKey])-1 {
				link.frameCount = 0

				// finished attack animation move switch to stand animation
				if link.isAttacking(link.lastFrameType) {
					link.lastFrameType = STAND
				}
			} else {
				link.frameCount++
			}
		} else {
			link.tick++
		}
	}

	link.currFrame = link.anims[frameKey][link.frameCount]
}

func (link *Link) getFrameKey(frameType int) string {
	switch frameType {
	case WALK_LEFT:
		return "walk_left"
	case WALK_RIGHT:
		return "walk_right"
	case WALK_DOWN:
		return "walk_down"
	case WALK_UP:
		return "walk_up"
	case STAND:
		return "stand"
	case ATTACK_UP:
		return "attack_up"
	case ATTACK_DOWN:
		return "attack_down"
	case ATTACK_LEFT:
		return "attack_left"
	case ATTACK_RIGHT:
		return "attack_right"
	}

	return "link_stand"
}

func (link *Link) Update(win *pixelgl.Window, world *World) int {
	frameType := STAND

	// don't move while attacking
	if !link.isAttacking(link.lastFrameType) {
		var newPos pixel.Vec
		var bouncePos pixel.Vec
		newPos, bouncePos, frameType = link.trackMovement(win)

		worldType := link.handleDoors(newPos, world.doors)
		if worldType != CURRENT {
			return worldType
		}

		link.Pos = link.handleCollisions(newPos, world, frameType, bouncePos)
	} else {
		frameType = link.lastFrameType
	}

	link.setCurrentFrame(frameType)
	return CURRENT
}

func (link *Link) handleCollisions(newPos pixel.Vec, world *World, frameType int, bouncePos pixel.Vec) pixel.Vec {
	validPosition := link.handleObstacleCollisions(newPos, world.objects)
	if validPosition {
		validPosition, newPos = link.handleEnemyCollisions(newPos, world.enemies, frameType, bouncePos)

		if validPosition {
			return newPos
		}
	}

	return link.Pos
}

func (link *Link) handleDoors(newPos pixel.Vec, doors []*Door) int {
	linkBounds := utils.GetBounds(newPos, pixel.R(0, 0, 30, 30))
	for _, d := range doors {
		if utils.Overlap(d.bounds, linkBounds) {
			return d.target
		}
	}

	return CURRENT
}

func (link *Link) handleObstacleCollisions(newPos pixel.Vec, objects []*Object) bool {
	linkBounds := utils.GetBounds(newPos, pixel.R(0, 0, 30, 30))
	for _, o := range objects {
		if o.blocking && utils.Overlap(o.bounds, linkBounds) {
			return false
		}
	}

	screenBounds := pixel.R(60, 60, 1024-60, 768-60)
	if !utils.Overlap(linkBounds, screenBounds) {
		return false
	}

	return true
}

func (link *Link) handleEnemyCollisions(newPos pixel.Vec, enemies []*Enemy, frameType int, bouncePos pixel.Vec) (bool, pixel.Vec) {
	linkBounds := utils.GetBounds(newPos, pixel.R(0, 0, 30, 30))
	linkAttackBounds := utils.GetBounds(newPos, pixel.R(-15, -15, 45, 45))

	for _, e := range enemies {
		if !e.isDead && !e.isDying {
			enemyBounds := utils.GetBounds(e.loc, e.size)
			if link.isAttacking(frameType) && utils.Overlap(enemyBounds, linkAttackBounds) {
				e.frameCount = 0
				e.tick = 0
				e.animCount = 0
				e.isDying = true
				return false, newPos
			} else if utils.Overlap(enemyBounds, linkBounds) {
				return true, bouncePos
			}
		}
	}

	return true, newPos
}

func (link *Link) trackMovement(win *pixelgl.Window) (pixel.Vec, pixel.Vec, int) {
	relPos := pixel.ZV
	newPos := link.Pos
	bouncePos := link.Pos
	frameType := STAND
	actionFrameType := ATTACK_UP

	if win.Pressed(pixelgl.KeyLeft) {
		newPos.X--
		relPos.X--
		bouncePos.X += 17
		frameType = WALK_LEFT
		actionFrameType = ATTACK_LEFT
	}
	if win.Pressed(pixelgl.KeyRight) {
		newPos.X++
		relPos.X++
		bouncePos.X -= 17
		frameType = WALK_RIGHT
		actionFrameType = ATTACK_RIGHT
	}
	if win.Pressed(pixelgl.KeyUp) {
		newPos.Y++
		relPos.Y++
		bouncePos.Y -= 17
		frameType = WALK_UP
		actionFrameType = ATTACK_UP
	}
	if win.Pressed(pixelgl.KeyDown) {
		newPos.Y--
		relPos.Y--
		bouncePos.Y += 17
		frameType = WALK_DOWN
		actionFrameType = ATTACK_DOWN
	}
	if relPos.X == 0 && relPos.Y == 0 {
		frameType = STAND
	}
	if win.JustPressed(pixelgl.KeySpace) {
		frameType = actionFrameType
	}

	return newPos, bouncePos, frameType
}

func (link *Link) isAttacking(frameType int) bool {
	return frameType == ATTACK_UP || frameType == ATTACK_DOWN || frameType == ATTACK_LEFT || frameType == ATTACK_RIGHT
}

func (link *Link) Draw(win *pixelgl.Window) {
	sprite := pixel.NewSprite(link.sheet, link.currFrame)
	sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(link.Pos))
}
