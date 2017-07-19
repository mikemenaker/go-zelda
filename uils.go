package main

import (
	"encoding/csv"
	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"image"
	"io"
	"os"
	"strconv"
)

func getBounds(location pixel.Vec, size pixel.Rect) pixel.Rect {
	return pixel.R(location.X-size.Max.X, location.Y-size.Max.Y, location.X+size.Max.X, location.Y+size.Max.Y)
}

func overlap(obj1 pixel.Rect, obj2 pixel.Rect) bool {
	empty := pixel.R(0, 0, 0, 0)
	return obj1.Intersect(obj2) != empty
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
