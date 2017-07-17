package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

func NewLink() *Link {
	link := new(Link)
	sheet, err := loadPicture("images/sprites/link.png")
	if err != nil {
		panic(err)
	}

	link.sheet = sheet
	link.anims = make(map[string][]pixel.Rect)
	link.anims["link_stand"] = []pixel.Rect{pixel.R(30, 360, 60, 390)}

	link.anims["link_down"] = []pixel.Rect{pixel.R(30, 330, 60, 360),
		pixel.R(60, 330, 90, 360),
		pixel.R(90, 330, 120, 360),
		pixel.R(120, 330, 150, 360)}

	link.anims["link_up"] = []pixel.Rect{pixel.R(0, 240, 30, 270),
		pixel.R(30, 240, 60, 270),
		pixel.R(60, 240, 90, 270),
		pixel.R(90, 240, 120, 270)}

	link.anims["link_left"] = []pixel.Rect{pixel.R(240, 360, 262, 390),
		pixel.R(262, 360, 287, 390),
		pixel.R(287, 360, 312, 390),
		pixel.R(310, 360, 337, 390)}

	link.anims["link_right"] = []pixel.Rect{pixel.R(240, 240, 270, 270),
		pixel.R(270, 240, 300, 270),
		pixel.R(300, 240, 330, 270),
		pixel.R(330, 240, 360, 270)}

	link.pos = pixel.ZV
	link.currFrame = pixel.R(0, 0, 0, 0)

	link.frameCount = 0
	link.lastFrameType = STAND
	link.tick = 0
	return link
}

const (
	LEFT = iota + 1
	RIGHT
	DOWN
	UP
	STAND
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
	case LEFT:
		return "link_left"
	case RIGHT:
		return "link_right"
	case DOWN:
		return "link_down"
	case UP:
		return "link_up"
	case STAND:
		return "link_stand"
	}
	return "link_stand"
}

func (link *Link) update(win *pixelgl.Window, objBound pixel.Rect, objPos pixel.Vec) {
	//fmt.Println(link.pos)





	objFullBound := pixel.R(objPos.X - objBound.Max.X / 2, objPos.Y + objBound.Max.Y / 2, objPos.X + objBound.Max.X / 2, objPos.Y - objBound.Max.Y / 2)

	frameType := STAND
	relPos := pixel.ZV
	if win.Pressed(pixelgl.KeyLeft) {
		link.pos.X--
		relPos.X--
		frameType = LEFT
	}

	//Min:Vec(429, 279) Max:Vec(625, 489)
	if win.Pressed(pixelgl.KeyRight) {
		tempPos := link.pos
		tempPos.X++
		fmt.Println("------")
		//fmt.Println(objFullBound)
		fmt.Println(tempPos)
		//fmt.Println(objBound)
		//fmt.Println(objPos)
		if (tempPos.X > objFullBound.Min.X && tempPos.X < objFullBound.Max.X) &&
			(tempPos.Y > objFullBound.Min.Y && tempPos.Y < objFullBound.Max.Y) {
			fmt.Println("collision")
		} else {
			link.pos.X++
		}
		relPos.X++
		frameType = RIGHT
	}
	if win.Pressed(pixelgl.KeyUp) {
		link.pos.Y++
		relPos.Y++
		frameType = UP
	}
	if win.Pressed(pixelgl.KeyDown) {
		link.pos.Y--
		relPos.Y--
		frameType = DOWN
	}
	if relPos.X == 0 && relPos.Y == 0 {
		frameType = STAND
	}

	link.setCurrentFrame(frameType)
}

func (link *Link) draw(win *pixelgl.Window) {
	sprite := pixel.NewSprite(link.sheet, link.currFrame)
	sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.5).Moved(link.pos))
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
