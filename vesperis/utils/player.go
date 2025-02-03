package utils

import (
	"github.com/team-vesperis/vesperis-proxy/vesperis/vanish"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/uuid"
)

func UUIDFromPlayerId(playerId string) uuid.UUID {
	UUID, _ := uuid.Parse(playerId)
	return UUID
}

func PlayerFromPlayerId(playerId string) proxy.Player {
	UUID := UUIDFromPlayerId(playerId)
	return p.Player(UUID)
}

func PlayerNameFromPlayerId(playerId string) string {
	UUID := UUIDFromPlayerId(playerId)
	return p.Player(UUID).Username()
}

func GetPlayerCount() int {
	var number int

	for _, player := range p.Players() {
		if !vanish.IsPlayerVanished(player) {
			number++
		}
	}

	return number
}
