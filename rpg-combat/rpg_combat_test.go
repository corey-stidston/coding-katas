package main

import (
	"errors"
	"testing"
)

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

	err := player1.dealDamage(player1, 100)

	if !errors.Is(err, ErrPlayersCannotDealDamageToThemselves) {
		t.Error("Expected error when player deals damage to themselves")
	}
}

func TestPlayerCannotTakeDamageIfItsAlreadyDead(t *testing.T) {
	player1 := Player()
	player2 := Player()

	player1.dealDamage(player2, 99999)

	err := player1.dealDamage(player2, 100)

	if !errors.Is(err, ErrDeadPlayersCannotBeDamaged) {
		t.Error("Expected error when dead player is damaged")
	}
}

func TestPlayerDeath(t *testing.T) {
	player1 := Player()
	player2 := Player()

	player1.dealDamage(player2, 9999)

	if player2.isAlive() {
		t.Error("Expected the player to be dead after receiving damage")
	}

	if player2.health != 0 {
		t.Error("Expected the dead player's health to be zero'")
	}
}

func TestHeal(t *testing.T) {
	player1 := Player()
	player2 := Player()

	healingAmount := 10
	startingHealth := player2.health
	player1.heal(player2, healingAmount)

	if player2.health != startingHealth+healingAmount {
		t.Errorf("Expected the player to be healed by %d", healingAmount)
	}
}

func TestPlayerCannotHealThemselves(t *testing.T) {
	player1 := Player()

	err := player1.heal(player1, 100)

	if !errors.Is(err, ErrPlayersCannotHealThemselves) {
		t.Error("Expected the player to have the same health as they cannot heal themselves")
	}
}

func TestPlayerMaximumHealth(t *testing.T) {

	tests := []struct {
		level int
		expectedHealth int
	}{
		{
			level: 1,
			expectedHealth: 1000,
		},
		{
			level: 5,
			expectedHealth: 1000,
		},
		{
			level: 6,
			expectedHealth: 1500,
		},
	}

	for _, test := range tests {
		player := player{
			level: test.level,
		}

		if player.getMaxHealth() != test.expectedHealth {
			t.Errorf("Expected the player at level %d to have a maximum health of %d, not %d", test.level, test.expectedHealth, player.getMaxHealth())
		}
	}
}

func TestAnAttacker5LevelsBelowTheTargetHasDamageReducedByHalf(t *testing.T) {
	player1 := Player()
	player2 := player{
		health: 1000,
		level: 6,
	}

	damageAmount := 100
	expectedHealth := player2.health - damageAmount / 2
	player1.dealDamage(&player2, damageAmount)

	if player2.health != expectedHealth {
		t.Errorf("Expected damage to be halved. Expected health: %d, Result: %d", expectedHealth, player2.health)
	}
}
