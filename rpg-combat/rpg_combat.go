package main

import "errors"

type player struct {
	health int
	level int
	factions map[*faction]bool
}

type faction struct {
	members map[*player]bool
}

func Player() *player {
	return &player{
		health: 1000,
		level: 1,
		factions: make(map[*faction]bool),
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

func (player *player) getFactions() []*faction {
	var factions []*faction
	for fac, isMemberOf := range player.factions {
		if isMemberOf {
			factions = append(factions, fac)
		}
	}
	return factions
}

func (player *player) joinFaction(faction *faction) {
	faction.members[player] = true
	player.factions[faction] = true
}

func (player *player) leaveFaction(faction *faction) {
	delete(faction.members, player)
	delete(player.factions, faction)
}

func Faction() *faction {
	return &faction{
		members: make(map[*player]bool),
	}
}

func (faction *faction) getMembers() []*player {
	var members []*player
	for mem, isMemberOf:= range faction.members {
		if isMemberOf {
			members = append(members, mem)
		}
	}
	return members
}

var ErrPlayersCannotDealDamageToThemselves = errors.New("players cannot deal damage to themselves")
var ErrDeadPlayersCannotBeDamaged = errors.New("dead players cannot be damaged")
var ErrPlayersCannotHealThemselves = errors.New("players cannot heal themselves")

func main() {
}
