package vanish

import (
	"github.com/team-vesperis/vesperis-proxy/vesperis/database"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func IsPlayerVanished(player proxy.Player) bool {
	playerId := player.ID().String()

	vanished, ok := database.GetPlayerDataField(playerId, "vanished").(bool)
	if !ok {
		vanished = false
		SetPlayerVanished(player, false)
	}

	return vanished
}

func SetPlayerVanished(player proxy.Player, vanished bool) {
	playerId := player.ID().String()

	database.SetPlayerDataField(playerId, "vanished", vanished)
}
