package main

import (
	_ "image/png"

	"encoding/csv"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"image"
	"io"
	"os"
	"strconv"
	"fmt"
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

func loadAnimationSheet(sheetPath, descPath string) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {
	// total hack, nicely format the error at the end, so I don't have to type it every time
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "error loading animation sheet")
		}
	}()

	// open and load the spritesheet
	sheetFile, err := os.Open(sheetPath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	descFile, err := os.Open(descPath)
	if err != nil {
		return nil, nil, err
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)

	// load the animation information, name and interval inside the spritesheet
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		name := anim[0]
		minx, _ := strconv.Atoi(anim[1])
		miny, _ := strconv.Atoi(anim[2])
		maxx, _ := strconv.Atoi(anim[3])
		maxy, _ := strconv.Atoi(anim[4])
		newFrame := pixel.R(float64(minx), float64(miny), float64(maxx), float64(maxy))

		if frames, ok := anims[name]; ok {
			frames = append(frames, newFrame)
			anims[name] = frames
		} else {
			anims[name] = []pixel.Rect{newFrame}
		}
	}

	return sheet, anims, nil
}

func NewLink() *Link {
	link := new(Link)

	sheet, anims, err := loadAnimationSheet("images/sprites/link.png", "images/sprites/sheet.csv")
	if err != nil {
		panic(err)
	}
	link.sheet = sheet
	link.anims = anims

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

func (link *Link) update(win *pixelgl.Window, objects []*Object, enemies []*Enemy) {
	frameType := STAND

	if !link.isAttacking(link.lastFrameType) {
		relPos := pixel.ZV
		newPos := link.pos
		bouncePos := link.pos
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

		overlapped := false
		linkBounds := getBounds(newPos, pixel.R(0, 0, 30, 30))
		linkAttackBounds := getBounds(newPos, pixel.R(-15, -15, 45, 45))

		for _, o := range objects {
			if overlap(o.bounds, linkBounds) {
				overlapped = true
				break
			}
		}

		for _, e := range enemies {
			if link.isAttacking(frameType) {
				if overlap(e.bounds, linkAttackBounds) {
					fmt.Println("collided and attacking")
					overlapped = true
				}
			} else {
				if overlap(e.bounds, linkBounds) {
					fmt.Println("collision no attack")
					newPos = bouncePos
				}
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
func (link *Link) isAttacking(frameType int) bool {
	return frameType == ATTACK_UP || frameType == ATTACK_DOWN || frameType == ATTACK_LEFT || frameType == ATTACK_RIGHT
}

func (link *Link) draw(win *pixelgl.Window) {
	sprite := pixel.NewSprite(link.sheet, link.currFrame)
	sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(link.pos))
}
