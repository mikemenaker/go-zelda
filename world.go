package main

import "github.com/faiface/pixel"

const (
	OVERWORLD = iota + 1
	CAVE
)

func createWorld() ([]*Object, []*Enemy) {
	var objects []*Object
	objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384), true))
	objects = append(objects, NewObject("images/tree.png", pixel.V(600, 384), true))

	objects = append(objects, NewObject("images/grass1.png", pixel.V(220, 670), false))
	objects = append(objects, NewObject("images/grass1.png", pixel.V(790, 220), false))
	objects = append(objects, NewObject("images/grass2.png", pixel.V(220, 220), false))
	objects = append(objects, NewObject("images/grass2.png", pixel.V(920, 620), false))
	//objects = append(objects, NewObject("images/grass3.png", pixel.V(20, 320), false))
	//objects = append(objects, NewObject("images/grass3.png", pixel.V(420, 20), false))

	var enemies []*Enemy
	enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
	enemies = append(enemies, NewEnemy(pixel.V(680, 600)))

	return objects, enemies
}
