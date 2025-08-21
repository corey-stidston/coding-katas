package main

type player struct {
	health int
}

func Player() *player {
	return &player{
		health: 1000,
	}
}

func main() {
}
