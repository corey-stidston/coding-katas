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
	expected := 900
	player := Player()

	player.damage(100)

	if player.health != expected {
		t.Errorf("Expected the player's health to be %d but was %d", expected, player.health)
	}
}

func TestPlayerDeath(t *testing.T) {
	player := Player()

	player.damage(1001)

	if player.isAlive() {
		t.Error("Expected the player to be dead after receiving damage")
	}
}
