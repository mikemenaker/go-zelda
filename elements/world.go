package elements

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type World struct {
	objects []*Object
	enemies []*Enemy
	doors   []*Door
	brColor color.Color
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
		return createOverworld()
	} else if worldType == CAVE {
		return createCave()
	} else if worldType == CASTLE {
		return createCastle()
	}

	return new(World)
}

func createOverworld() *World {
	world := new(World)
	var objects []*Object
	objects = append(objects, NewObject("images/overworld/tree.png", pixel.V(200, 384), true))
	objects = append(objects, NewObject("images/overworld/tree.png", pixel.V(600, 384), true))

	objects = append(objects, NewObject("images/overworld/grass1.png", pixel.V(220, 670), false))
	objects = append(objects, NewObject("images/overworld/grass1.png", pixel.V(245, 670), false))
	objects = append(objects, NewObject("images/overworld/grass1.png", pixel.V(790, 220), false))
	objects = append(objects, NewObject("images/overworld/grass1.png", pixel.V(790, 245), false))
	objects = append(objects, NewObject("images/overworld/grass2.png", pixel.V(220, 220), false))
	objects = append(objects, NewObject("images/overworld/grass2.png", pixel.V(220, 245), false))
	objects = append(objects, NewObject("images/overworld/grass2.png", pixel.V(920, 620), false))
	objects = append(objects, NewObject("images/overworld/grass2.png", pixel.V(945, 620), false))
	objects = append(objects, NewObject("images/overworld/dirt_patch.png", pixel.V(775, 520), false))
	objects = append(objects, NewObject("images/overworld/dirt_patch.png", pixel.V(380, 150), false))
	world.objects = objects

	var enemies []*Enemy
	enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
	enemies = append(enemies, NewEnemy(pixel.V(680, 600)))
	world.enemies = enemies

	var doors []*Door
	doors = append(doors, NewDoor("images/overworld/cave_entrance.png", pixel.V(390, 670), CAVE))
	world.doors = doors

	world.brColor = color.RGBA{72, 152, 72, 1}

	world.LinkPos = pixel.V(0, 0)
	return world
}

func createCave() *World {
	world := new(World)
	var objects []*Object
	objects = append(objects, NewObject("images/cave/wall_right.png", pixel.V(980, 0), true))
	objects = append(objects, NewObject("images/cave/wall_right.png", pixel.V(980, 124), true))
	objects = append(objects, NewObject("images/cave/wall_right.png", pixel.V(980, 248), true))
	objects = append(objects, NewObject("images/cave/wall_right.png", pixel.V(980, 372), true))
	objects = append(objects, NewObject("images/cave/wall_right.png", pixel.V(980, 496), true))
	objects = append(objects, NewObject("images/cave/wall_right.png", pixel.V(980, 620), true))
	objects = append(objects, NewObject("images/cave/wall_left.png", pixel.V(24, 0), true))
	objects = append(objects, NewObject("images/cave/wall_left.png", pixel.V(24, 124), true))
	objects = append(objects, NewObject("images/cave/wall_left.png", pixel.V(24, 248), true))
	objects = append(objects, NewObject("images/cave/wall_left.png", pixel.V(24, 372), true))
	objects = append(objects, NewObject("images/cave/wall_left.png", pixel.V(24, 496), true))
	objects = append(objects, NewObject("images/cave/wall_left.png", pixel.V(24, 620), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(0, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(132, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(264, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(394, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(526, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(658, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(790, 730), true))
	objects = append(objects, NewObject("images/cave/wall_top.png", pixel.V(922, 730), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(0, 0), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(132, 24), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(264, 24), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(394, 24), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(526, 24), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(658, 24), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(790, 24), true))
	objects = append(objects, NewObject("images/cave/wall_bottom.png", pixel.V(922, 24), true))
	objects = append(objects, NewObject("images/cave/top_left.png", pixel.V(24, 730), true))
	objects = append(objects, NewObject("images/cave/top_right.png", pixel.V(972, 730), true))
	objects = append(objects, NewObject("images/cave/bottom_left.png", pixel.V(24, 24), true))
	objects = append(objects, NewObject("images/cave/bottom_right.png", pixel.V(972, 24), true))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(220, 570), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(245, 570), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(790, 220), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(790, 245), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(220, 220), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(220, 245), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(720, 620), false))
	objects = append(objects, NewObject("images/cave/floor_tile.png", pixel.V(745, 620), false))
	objects = append(objects, NewObject("images/cave/rock.png", pixel.V(550, 550), true))
	world.objects = objects
	var enemies []*Enemy
	enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
	enemies = append(enemies, NewEnemy(pixel.V(680, 600)))
	world.enemies = enemies
	var doors []*Door
	doors = append(doors, NewDoor("images/cave/exit.png", pixel.V(515, 710), OVERWORLD))
	world.doors = doors

	world.brColor = color.RGBA{40, 32, 32, 1}
	world.LinkPos = pixel.V(0, 0)
	return world
}

func createCastle() *World {
	world := new(World)
	var objects []*Object
	objects = append(objects, NewObject("images/overworld/tree.png", pixel.V(200, 384), true))
	world.objects = objects

	var enemies []*Enemy
	enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
	enemies = append(enemies, NewEnemy(pixel.V(680, 600)))
	world.enemies = enemies

	world.brColor = colornames.Skyblue

	world.LinkPos = pixel.V(0, 0)
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
