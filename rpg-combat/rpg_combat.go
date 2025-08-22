package main

type player struct {
	health int
}

func Player() *player {
	return &player{
		health: 1000,
	}
}

func (player *player) dealDamage(target *player, amount int) {
	if player == target {
		return // player cannot deal damage to itself
	}
	if !target.isAlive() {
		return // player cannot take damage if it is dead
	}

	if amount > target.health {
		target.health = 0 // player health cannot be less than zero
	} else {
		target.health -= amount
	}
}

func (player *player) heal(target *player, amount int) {

}

func (player *player) isAlive() bool {
	return player.health > 0
}

func main() {
}
