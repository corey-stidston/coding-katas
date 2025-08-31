package main

import (
	"errors"
	"testing"
)

func testSetup2PlayersInTheSameFaction() (*player, *player, *faction) {
	player1 := Player()
	player2 := Player()
	faction1 := Faction()

	player1.joinFaction(faction1)
	player2.joinFaction(faction1)

	return player1, player2, faction1
}

func testSetup2PlayersInSeparateFactions() (*player, *player, *faction, *faction) {
	player1 := Player()
	player2 := Player()
	faction1 := Faction()
	faction2 := Faction()

	player1.joinFaction(faction1)
	player2.joinFaction(faction2)

	return player1, player2, faction1, faction2
}

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

func TestPlayerCannotDealDamageToAllies(t *testing.T) {
	var player1, player2, _ = testSetup2PlayersInTheSameFaction()

	err := player1.dealDamage(player2, 100)

	if !errors.Is(err, ErrPlayersCannotDamageAllies) {
		t.Error("Expected error when player deals damage to an ally")
	}
}

func TestPlayerDealingDamageToNonAlly(t *testing.T) {
	player1, player2, _, _ := testSetup2PlayersInSeparateFactions()

	damage := 100
	startingHealth := player2.health
	expectedHealth := startingHealth - damage
	player1.dealDamage(player2, damage)

	if player2.health != expectedHealth {
		t.Errorf("Expected player 2 health to be %d, but was %d", expectedHealth, player2.health)
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

func TestPlayersThatAreAlliesCanHeal(t *testing.T) {
	player1, player2, _ := testSetup2PlayersInTheSameFaction()

	healingAmount := 10
	startingHealth := player2.health
	player1.heal(player2, healingAmount)

	if player2.health != startingHealth+healingAmount {
		t.Errorf("Expected the player to be healed by %d", healingAmount)
	}
}

func TestPlayerCanOnlyBeHealedUpToMaximumHealth(t *testing.T) {
	player1, player2, _ := testSetup2PlayersInTheSameFaction()

	healingAmount := 1000
	startingHealth := player2.health
	player1.heal(player2, healingAmount)

	if player2.health != startingHealth {
		t.Errorf("Expected the player to be healed up to their maximum health %d but was %d", startingHealth, player2.health)
	}
}

func TestPlayersThatAreNotAlliesCannotHeal(t *testing.T) {
	player1, player2, _, _ := testSetup2PlayersInSeparateFactions()

	err := player1.heal(player2, 10)

	if !errors.Is(err, ErrPlayersCanOnlyHealAllies) {
		t.Error("Expected an error that players can only heal allies")
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
		level          int
		expectedHealth int
	}{
		{
			level:          1,
			expectedHealth: 1000,
		},
		{
			level:          5,
			expectedHealth: 1000,
		},
		{
			level:          6,
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
		level:  6,
	}

	damageAmount := 100
	expectedHealth := player2.health - damageAmount/2
	player1.dealDamage(&player2, damageAmount)

	if player2.health != expectedHealth {
		t.Errorf("Expected damage to be halved. Expected health: %d, Result: %d", expectedHealth, player2.health)
	}
}

func TestAnAttacker5LevelsAboveTheTargetHasDamageIncreasedBy50pc(t *testing.T) {
	player1 := player{
		health: 1000,
		level:  6,
	}
	player2 := Player()

	damageAmount := 100
	expectedHealth := player2.health - damageAmount - damageAmount/2
	player1.dealDamage(player2, damageAmount)

	if player2.health != expectedHealth {
		t.Errorf("Expected damage to be increased by 50pc. Expected health: %d, Result: %d", expectedHealth, player2.health)
	}
}

func TestNewPlayersDoNotBelongToAFaction(t *testing.T) {
	player := Player()

	if len(player.factions) != 0 {
		t.Error("Expected that a new player does not belong to a faction")
	}
}

func TestPlayerJoiningFaction(t *testing.T) {
	faction := Faction()
	player := Player()

	player.joinFaction(faction)

	if len(player.getFactions()) != 1 || player.getFactions()[0] != faction {
		t.Error("Expected the player to belong to the faction")
	}

	if len(faction.getMembers()) != 1 || faction.getMembers()[0] != player {
		t.Error("Expected the faction to have the player as a member")
	}
}

func TestPlayerLeavingFaction(t *testing.T) {
	faction := Faction()
	player := Player()

	player.joinFaction(faction)
	player.leaveFaction(faction)

	if len(player.getFactions()) != 0 {
		t.Error("Expected the player to no longer be a member of the faction")
	}

	if len(faction.members) != 0 {
		t.Error("Expected the faction to have no members")
	}
}

func TestAPlayerCannotJoinFactionMultipleTimes(t *testing.T) {
	faction := Faction()
	player := Player()

	player.joinFaction(faction)
	player.joinFaction(faction)

	if len(player.getFactions()) != 1 {
		t.Error("Expected the player to belong in the faction exactly once")
	}

	if len(faction.members) != 1 {
		t.Error("Expected the faction to have exactly one member")
	}
}
