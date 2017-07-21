package main

import "github.com/faiface/pixel"

func createWorld() ([]*Object, []*Enemy) {
	var objects []*Object
	objects = append(objects, NewObject("images/tree.png", pixel.V(200, 384)))
	objects = append(objects, NewObject("images/tree.png", pixel.V(600, 384)))

	var enemies []*Enemy
	enemies = append(enemies, NewEnemy(pixel.V(440, 384)))
	enemies = append(enemies, NewEnemy(pixel.V(680, 600)))

	return objects, enemies
}
