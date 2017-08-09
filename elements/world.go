package elements

import (
	"encoding/csv"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"io"
	"os"
	"strconv"
	"fmt"
)

type World struct {
	objects []*Object
	enemies []*Enemy
	doors   []*Door
	bgColor color.Color
	LinkPos pixel.Vec
}

const (
	CURRENT = iota + 1
	OVERWORLD
	CAVE
	CASTLE
)

func CreateWorld(worldType int) *World {
	if worldType == OVERWORLD {
		return readWorld("elements/overworld.csv")
	} else if worldType == CAVE {
		return readWorld("elements/cave.csv")
	} else if worldType == CASTLE {
		return readWorld("elements/castle.csv")
	}

	return new(World)
}

func (world *World) UpdateAndDraw(win *pixelgl.Window) {
	win.Clear(world.bgColor)

	for _, o := range world.objects {
		o.draw(win)
	}

	for _, d := range world.doors {
		d.draw(win)
	}

	for _, e := range world.enemies {
		if !e.isDead {
			e.update(win, world.objects, world.enemies)
			e.draw(win)
		}
	}

}

func readWorld(path string) *World {
	world := new(World)
	var objects []*Object
	var enemies []*Enemy
	var doors []*Door
	world.LinkPos = pixel.V(0, 0)

	descFile, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer descFile.Close()

	desc := csv.NewReader(descFile)
	desc.FieldsPerRecord = -1
	for {
		element, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}

		switch element[0] {
		case "object":
			x, _ := strconv.Atoi(element[2])
			y, _ := strconv.Atoi(element[3])
			blocking, _ := strconv.ParseBool(element[4])
			objects = append(objects, NewObject(element[1], pixel.V(float64(x), float64(y)), blocking))
			fmt.Println(pixel.V(float64(x), float64(y)))
		case "enemy":
			x, _ := strconv.Atoi(element[2])
			y, _ := strconv.Atoi(element[3])
			enemies = append(enemies, NewEnemy(pixel.V(float64(x), float64(y)), element[1]))
		case "door":
			x, _ := strconv.Atoi(element[2])
			y, _ := strconv.Atoi(element[3])
			target, _ := strconv.Atoi(element[4])
			doors = append(doors, NewDoor(element[1], pixel.V(float64(x), float64(y)), target))
		case "linkPos":
			x, _ := strconv.Atoi(element[1])
			y, _ := strconv.Atoi(element[2])
			world.LinkPos = pixel.V(float64(x), float64(y))
		case "bgColor":
			r, _ := strconv.Atoi(element[1])
			g, _ := strconv.Atoi(element[2])
			b, _ := strconv.Atoi(element[3])
			world.bgColor = color.RGBA{uint8(r), uint8(g), uint8(b), 1}
		}
	}

	world.objects = objects
	world.enemies = enemies
	world.doors = doors
	return world

}
