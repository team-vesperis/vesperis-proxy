package vanish

import "github.com/team-vesperis/vesperis-proxy/database"

func IsPlayerVanished(playerId string) bool {
	data := database.GetPlayerData(playerId)
	vanished, ok := data["vanished"].(bool)
	if !ok {
		vanished = false
		SetPlayerVanished(playerId, false)
	}

	return vanished
}

func SetPlayerVanished(playerId string, vanished bool) {
	data := database.GetPlayerData(playerId)
	data["vanished"] = vanished
	database.SavePlayerData(playerId, data)
}
