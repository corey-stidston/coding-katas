package main

type healingMagicalObject struct {
	health int
}

type magicalWeapon struct {
	health       int
	damagePoints int
}

func HealingMagicalObject(health int) *healingMagicalObject {
	return &healingMagicalObject{
		health: health,
	}
}

func (hmo *healingMagicalObject) isDestroyed() bool {
	return isDestroyed(hmo.health)
}

func MagicalWeapon(health int, damagePoints int) *magicalWeapon {
	return &magicalWeapon{
		health:       health,
		damagePoints: damagePoints,
	}
}

func (weapon *magicalWeapon) isDestroyed() bool {
	return isDestroyed(weapon.health)
}

func isDestroyed(health int) bool {
	return health <= 0
}
