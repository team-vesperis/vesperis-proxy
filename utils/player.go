package utils

import (
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
