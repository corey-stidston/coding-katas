package main

type faction struct {
	members map[*player]playerFactionLink
}

type playerFactionLink struct{}

func Faction() *faction {
	return &faction{
		members: make(map[*player]playerFactionLink),
	}
}

func (faction *faction) getMembers() []*player {
	var members []*player
	for mem := range faction.members {
		members = append(members, mem)
	}
	return members
}
