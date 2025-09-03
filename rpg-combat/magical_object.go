package main

type healingMagicalObject struct {
	health int
}

func HealingMagicalObject(health int) *healingMagicalObject {
	return &healingMagicalObject{
		health: health,
	}
}

func (hmo *healingMagicalObject) isDestroyed() bool {
	return hmo.health <= 0
}