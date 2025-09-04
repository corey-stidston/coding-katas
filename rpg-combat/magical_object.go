package main

type healingMagicalObject struct {
	health int
}

type magicalWeapon struct {
	health int
	hitPoints int
}

func HealingMagicalObject(health int) *healingMagicalObject {
	return &healingMagicalObject{
		health: health,
	}
}

func (hmo *healingMagicalObject) isDestroyed() bool {
	return hmo.health <= 0
}

func MagicalWeapon(health int, hitPoints int) *magicalWeapon {
	return &magicalWeapon{
		health: health,
		hitPoints: hitPoints,
	}
}
