package main

import (
	"github.com/faiface/pixel"
	"image"
	"os"
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
