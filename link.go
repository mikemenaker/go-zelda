package main

import (
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Link struct {
	sheet         pixel.Picture
	anims         map[string][]pixel.Rect
	pos           pixel.Vec
	currFrame     pixel.Rect
	lastFrameType int
	frameCount    int
	tick          int
}

func NewLink() *Link {
	link := new(Link)
	sheet, err := loadPicture("images/sprites/link.png")
	if err != nil {
		panic(err)
	}

	link.sheet = sheet
	link.anims = make(map[string][]pixel.Rect)
	link.anims["stand"] = []pixel.Rect{pixel.R(30, 360, 60, 390)}

	link.anims["walk_down"] = []pixel.Rect{pixel.R(30, 330, 60, 360),
		pixel.R(60, 330, 90, 360),
		pixel.R(90, 330, 120, 360),
		pixel.R(120, 330, 150, 360)}

	link.anims["walk_up"] = []pixel.Rect{pixel.R(0, 240, 30, 270),
		pixel.R(30, 240, 60, 270),
		pixel.R(60, 240, 90, 270),
		pixel.R(90, 240, 120, 270)}

	link.anims["walk_left"] = []pixel.Rect{pixel.R(240, 360, 262, 390),
		pixel.R(262, 360, 287, 390),
		pixel.R(287, 360, 312, 390),
		pixel.R(310, 360, 337, 390)}

	link.anims["walk_right"] = []pixel.Rect{pixel.R(240, 240, 270, 270),
		pixel.R(270, 240, 300, 270),
		pixel.R(300, 240, 330, 270),
		pixel.R(330, 240, 360, 270)}

	link.anims["attack_down"] = []pixel.Rect{pixel.R(0, 270, 30, 300),
		pixel.R(30, 270, 60, 300),
		pixel.R(60, 270, 90, 303),
		pixel.R(90, 270, 115, 303),
		pixel.R(115, 270, 145, 303),
		pixel.R(145, 270, 180, 303)}

	link.anims["attack_up"] = []pixel.Rect{pixel.R(0, 180, 25, 210),
		pixel.R(29, 172, 55, 213),
		pixel.R(60, 172, 85, 214),
		pixel.R(87, 174, 114, 213),
		pixel.R(115, 180, 145, 210)}

	link.anims["attack_left"] = []pixel.Rect{pixel.R(240, 270, 268, 300),
		pixel.R(270, 270, 295, 300),
		pixel.R(295, 270, 325, 300),
		pixel.R(326, 270, 358, 300),
		pixel.R(358, 270, 390, 303),
		pixel.R(358, 270, 390, 303),
		pixel.R(358, 270, 390, 303)}

	link.anims["attack_right"] = []pixel.Rect{pixel.R(240, 180, 265, 210),
		pixel.R(265, 180, 295, 210),
		pixel.R(295, 180, 327, 210),
		pixel.R(327, 180, 359, 210),
		pixel.R(359, 183, 392, 213)}

	link.pos = pixel.V(130, 130)
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
	frameKey := getFrameKey(frameType)

	if frameType != link.lastFrameType {
		link.lastFrameType = frameType
		link.frameCount = 0
		link.tick = 0
	} else {
		if link.tick == 3 {
			link.tick = 0
			if link.frameCount == len(link.anims[frameKey])-1 {
				link.frameCount = 0
				if link.lastFrameType == ATTACK_UP || link.lastFrameType == ATTACK_DOWN || link.lastFrameType == ATTACK_LEFT || link.lastFrameType == ATTACK_RIGHT {
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

func getFrameKey(frameType int) string {
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

func (link *Link) update(win *pixelgl.Window, objects []*Object) {
	frameType := STAND

	if link.lastFrameType != ATTACK_UP && link.lastFrameType != ATTACK_DOWN && link.lastFrameType != ATTACK_LEFT && link.lastFrameType != ATTACK_RIGHT {
		relPos := pixel.ZV
		newPos := link.pos
		actionFrameType := ATTACK_UP
		if win.Pressed(pixelgl.KeyLeft) {
			newPos.X--
			relPos.X--
			frameType = WALK_LEFT
			actionFrameType = ATTACK_LEFT
		}
		if win.Pressed(pixelgl.KeyRight) {
			newPos.X++
			relPos.X++
			frameType = WALK_RIGHT
			actionFrameType = ATTACK_RIGHT
		}
		if win.Pressed(pixelgl.KeyUp) {
			newPos.Y++
			relPos.Y++
			frameType = WALK_UP
			actionFrameType = ATTACK_UP
		}
		if win.Pressed(pixelgl.KeyDown) {
			newPos.Y--
			relPos.Y--
			frameType = WALK_DOWN
			actionFrameType = ATTACK_DOWN
		}
		if relPos.X == 0 && relPos.Y == 0 {
			frameType = STAND
		}

		if win.JustPressed(pixelgl.KeySpace) {
			frameType = actionFrameType
		}

		overlapped := false
		linkBounds := getBounds(newPos, pixel.R(0, 0, 30, 30))

		for _, o := range objects {
			if overlap(o.bounds, linkBounds) {
				overlapped = true
				break
			}
		}

		if !overlapped {
			link.pos = newPos
		}
	} else {
		frameType = link.lastFrameType
	}

	link.setCurrentFrame(frameType)
}

func (link *Link) draw(win *pixelgl.Window) {
	sprite := pixel.NewSprite(link.sheet, link.currFrame)
	sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(link.pos))
}
