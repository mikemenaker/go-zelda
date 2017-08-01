package main

import (
	"github.com/faiface/pixel"
	"image/color"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type World struct {
	objects []*Object
	enemies []*Enemy
	doors []*Door
	brColor color.Color
	linkPos pixel.Vec
}

const (
	CURRENT = iota + 1
	OVERWORLD
	CAVE
	CASTLE
)

func createWorld(worldType int) *World {
	world := new(World)
	if worldType == OVERWORLD {
		var objects []*Object
		objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384), true))
		objects = append(objects, NewObject("images/tree.png", pixel.V(600, 384), true))

		objects = append(objects, NewObject("images/grass1.png", pixel.V(220, 670), false))
		objects = append(objects, NewObject("images/grass1.png", pixel.V(245, 670), false))
		objects = append(objects, NewObject("images/grass1.png", pixel.V(790, 220), false))
		objects = append(objects, NewObject("images/grass1.png", pixel.V(790, 245), false))
		objects = append(objects, NewObject("images/grass2.png", pixel.V(220, 220), false))
		objects = append(objects, NewObject("images/grass2.png", pixel.V(220, 245), false))
		objects = append(objects, NewObject("images/grass2.png", pixel.V(920, 620), false))
		objects = append(objects, NewObject("images/grass2.png", pixel.V(945, 620), false))
		world.objects = objects

		var enemies []*Enemy
		enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
		enemies = append(enemies, NewEnemy(pixel.V(680, 600)))
		world.enemies = enemies

		var doors []*Door
		doors = append(doors, NewDoor("images/cave_entrance.png", pixel.V(390, 670), CAVE))
		world.doors = doors

		world.brColor = color.RGBA{72, 152, 72, 1}

		world.linkPos = pixel.V(0, 0)
	} else if worldType == CAVE {
		var objects []*Object
		objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384), true))
		world.objects = objects

		var enemies []*Enemy
		enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
		enemies = append(enemies, NewEnemy(pixel.V(680, 600)))
		world.enemies = enemies

		world.brColor = color.Black

		world.linkPos = pixel.V(0, 0)
	} else if worldType == CAVE {
		var objects []*Object
		objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384), true))
		world.objects = objects

		var enemies []*Enemy
		enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
		enemies = append(enemies, NewEnemy(pixel.V(680, 600)))
		world.enemies = enemies

		world.brColor = colornames.Skyblue

		world.linkPos = pixel.V(0, 0)
	}

	return world
}

func (world *World) UpdateAndDraw(win *pixelgl.Window) {
	win.Clear(world.brColor)

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