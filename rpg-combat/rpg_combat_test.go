package main

import "testing"

func TestStartingHealth(t *testing.T) {
	player := Player()

	if player.health != 1000 {
		t.Error("Expected the player's starting health to be 1000")
	}
}

func TestStartingAlive(t *testing.T) {
	player := Player()

	if !player.isAlive() {
		t.Error("Expected the player to start alive")
	}
}

func TestDealingDamage(t *testing.T) {
    player1 := Player()
    player2 := Player()

	expected := 900
	
    player1.dealDamage(player2, 100)

	if player2.health != expected {
		t.Errorf("Expected the player's health to be %d but was %d", expected, player2.health)
	}
}

func TestPlayerCannotDealDamageToItself(t *testing.T) {
    player1 := Player()

    startingHealth := player1.health

    player1.dealDamage(player1, 100)

    if player1.health != startingHealth {
        t.Error("Expected player's health to be unchanged as player's cannot deal damanage to themselves")
    }
}

func TestPlayerDeath(t *testing.T) {
	player1 := Player()
    player2 := Player()

	player1.dealDamage(player2, 9999)

	if player2.isAlive() {
		t.Error("Expected the player to be dead after receiving damage")
	}
}
