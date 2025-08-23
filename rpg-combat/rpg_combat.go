package main

import "errors"

type player struct {
	health int
	level int
}

func Player() *player {
	return &player{
		health: 1000,
		level: 1,
	}
}

func (player *player) dealDamage(target *player, amount int) error {
	if player == target {
		return ErrPlayersCannotDealDamageToThemselves
	}
	if !target.isAlive() {
		return ErrDeadPlayersCannotBeDamaged
	}

	if target.level - player.level >= 5 {
		amount = amount / 2
	} else if player.level - target.level >= 5 {
		amount = amount + amount / 2
	}

	if amount > target.health {
		target.health = 0 // player health cannot be less than zero
	} else {
		target.health -= amount
	}
	return nil
}

func (player *player) heal(target *player, amount int) error {
	if player == target {
		return ErrPlayersCannotHealThemselves
	}

	target.health += amount
	return nil
}

func (player *player) isAlive() bool {
	return player.health > 0
}

func (player *player) getMaxHealth() int {
	if player.level < 6 {
		return 1000
	} else {
		return 1500
	}
}

var ErrPlayersCannotDealDamageToThemselves = errors.New("players cannot deal damage to themselves")
var ErrDeadPlayersCannotBeDamaged = errors.New("dead players cannot be damaged")
var ErrPlayersCannotHealThemselves = errors.New("players cannot heal themselves")

func main() {
}
