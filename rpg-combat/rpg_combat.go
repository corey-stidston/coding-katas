package main

type player struct {
	health int
}

func Player() *player {
	return &player{
		health: 1000,
	}
}

func (player *player) dealDamage(amount int) {
	if player.isAlive() {
		player.health -= amount
	}
}

func (player *player) isAlive() bool {
	return player.health > 0
}

func main() {
}
