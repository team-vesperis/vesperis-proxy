package vanish

import "github.com/team-vesperis/vesperis-proxy/database"

func IsPlayerVanished(playerId string) bool {
	vanished, ok := database.GetPlayerDataField(playerId, "vanished").(bool)
	if !ok {
		vanished = false
		SetPlayerVanished(playerId, false)
	}

	return vanished
}

func SetPlayerVanished(playerId string, vanished bool) {
	database.SetPlayerDataField(playerId, "vanished", vanished)
}
