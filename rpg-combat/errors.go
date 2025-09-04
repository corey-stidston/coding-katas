package main

import "errors"

var ErrPlayersCannotDamageAllies = errors.New("players cannot deal damage to allies")
var ErrPlayersCannotDealDamageToThemselves = errors.New("players cannot deal damage to themselves")
var ErrDeadPlayersCannotBeDamaged = errors.New("dead players cannot be damaged")
var ErrPlayersCannotHealThemselves = errors.New("players cannot heal themselves")
var ErrPlayersCanOnlyHealAllies = errors.New("players can only heal allies")
var ErrorCannotUseDestroyedMagicalObject = errors.New("cannot use destroyed magical object")
