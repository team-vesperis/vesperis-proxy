package permission

import (
	"github.com/team-vesperis/vesperis-proxy/database"
	"github.com/team-vesperis/vesperis-proxy/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

var validRoles = map[string]struct{}{
	"admin":     {},
	"builder":   {},
	"default":   {},
	"moderator": {},
}

var validRanks = map[string]struct{}{
	"champion": {},
	"default":  {},
	"elite":    {},
	"legend":   {},
}

func InitializePermissionManager(log *zap.SugaredLogger) {
	logger = log
	logger.Info("Initialized Permission Manager.")
}

func GetPlayerRole(playerId string) string {
	roleInterface := database.GetPlayerDataField(playerId, "role")
	role := "default"
	if roleInterface == nil {
		SetPlayerRole(playerId, "default")
	} else {
		role = roleInterface.(string)
	}

	_, valid := validRoles[role]
	if !valid {
		role = "default"
		SetPlayerRole(playerId, "default")
	}

	return role
}

func SetPlayerRole(playerId, role string) {
	_, valid := validRoles[role]
	if !valid {
		logger.Error("Invalid role: ", role)
		return
	}

	database.SetPlayerDataField(playerId, "role", role)

	playerName := utils.PlayerNameFromPlayerId(playerId)
	logger.Info("Changed permission role for " + playerName + " - " + playerId + " to " + role)
}

func IsPlayerPrivileged(playerId string) bool {
	role := GetPlayerRole(playerId)
	return role == "admin" || role == "builder" || role == "moderator"
}

func GetPlayerRank(playerId string) string {
	rankInterface := database.GetPlayerDataField(playerId, "rank")
	rank := "default"
	if rankInterface == nil {
		SetPlayerRank(playerId, "default")
	} else {
		rank = rankInterface.(string)
	}

	_, valid := validRanks[rank]
	if !valid {
		rank = "default"
		SetPlayerRank(playerId, "default")
	}

	return rank
}

func SetPlayerRank(playerId, rank string) {
	_, valid := validRanks[rank]
	if !valid {
		logger.Error("Invalid rank: ", rank)
		return
	}

	database.SetPlayerDataField(playerId, "rank", rank)

	playerName := utils.PlayerNameFromPlayerId(playerId)
	logger.Info("Changed permission rank for " + playerName + " - " + playerId + " to " + rank)
}
