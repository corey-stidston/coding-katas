package main

type player struct {
	health         int
	level          int
	factions       map[*faction]playerFactionLink
	totalDamageReceived int
}

func Player() *player {
	return &player{
		health:   1000,
		level:    1,
		factions: make(map[*faction]playerFactionLink),
		totalDamageReceived: 0,
	}
}

func (player *player) dealDamage(target *player, amount int) error {
	if player == target {
		return ErrPlayersCannotDealDamageToThemselves
	}
	if !target.isAlive() {
		return ErrDeadPlayersCannotBeDamaged
	}
	if player.isAnAllyTo(target) {
		return ErrPlayersCannotDamageAllies
	}

	if target.level-player.level >= 5 {
		amount = amount / 2
	} else if player.level-target.level >= 5 {
		amount = amount + amount/2
	}

	target.receiveDamage(amount)

	return nil
}

func (player *player) receiveDamage(amount int) {
	if amount > player.health {
		amount = player.health // player cannot receive more damage than their health
	}

	player.health -= amount
	player.totalDamageReceived += amount

	if player.health > 0 {
		if player.totalDamageReceived >= playerLevelTotalDamageReceivedMap[player.level + 1] {
			player.level++
		}
	}
}

var playerLevelTotalDamageReceivedMap = map[int]int{
	2: 1000,
	3: 3000,
	4: 6000,
	5: 10000,
	6: 15000,
	7: 21000,
	8: 28000,
	9: 36000,
	10: 45000,
}

func (player *player) heal(target *player, amount int) error {
	if player == target {
		return ErrPlayersCannotHealThemselves
	}
	if !player.isAnAllyTo(target) {
		return ErrPlayersCanOnlyHealAllies
	}

	healthDifference := target.getMaxHealth() - target.health
	if amount > healthDifference {
		amount = healthDifference
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
	for fac := range player.factions {
		factions = append(factions, fac)
	}
	return factions
}

func (player *player) joinFaction(faction *faction) {
	faction.members[player] = playerFactionLink{}
	player.factions[faction] = playerFactionLink{}
}

func (player *player) leaveFaction(faction *faction) {
	delete(faction.members, player)
	delete(player.factions, faction)
}

func (player *player) isAnAllyTo(target *player) bool {
	var allies bool
	for fac := range player.factions {
		_, ok := target.factions[fac]
		if ok {
			allies = true
			break
		}
	}
	return allies
}

func (player *player) healWithMagicalObject(hmo *healingMagicalObject, amount int) error {
	if hmo.isDestroyed() {
		return ErrorCannotUseDestroyedMagicalObject
	}

	if amount > hmo.health {
		amount = hmo.health
	}

	healthDifference := player.getMaxHealth() - player.health
	if amount > healthDifference {
		amount = healthDifference
	}

	player.health += amount
	hmo.health -= amount

	return nil
}

func (player *player) dealDamageWithWeapon(target *player, weapon *magicalWeapon) error {
	if weapon.isDestroyed() {
		return ErrorCannotUseDestroyedMagicalObject
	}

	err := player.dealDamage(target, weapon.damagePoints)

	if err != nil {
		return err
	}

	weapon.health--
	return nil
}
