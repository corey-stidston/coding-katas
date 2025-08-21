package main

import "testing"

func TestStartingHealth(t *testing.T) {
	player := Player()

	if player.health != 1000 {
		t.Error("Expected the player's starting health to be 1000")
	}
}

func TestDealingDamage(t *testing.T) {
    expected := 900
    player := Player()

    player.dealDamage(100)

    if player.health != expected {
        t.Errorf("Expected the player's health to be %d but was %d", expected, player.health)
    }
}

