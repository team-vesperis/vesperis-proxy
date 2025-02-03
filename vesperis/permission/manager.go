package permission

import (
	"github.com/team-vesperis/vesperis-proxy/vesperis/database"
	"go.minekube.com/gate/pkg/edition/java/proxy"
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

func GetPlayerRole(player proxy.Player) string {
	playerId := player.ID().String()

	roleInterface := database.GetPlayerDataField(playerId, "role")
	role := "default"
	if roleInterface == nil {
		SetPlayerRole(player, "default")
	} else {
		role = roleInterface.(string)
	}

	_, valid := validRoles[role]
	if !valid {
		role = "default"
		SetPlayerRole(player, "default")
	}

	return role
}

func SetPlayerRole(player proxy.Player, role string) {
	playerId := player.ID().String()

	_, valid := validRoles[role]
	if !valid {
		logger.Error("Invalid role: ", role)
		return
	}

	database.SetPlayerDataField(playerId, "role", role)
	logger.Info("Changed permission role for " + player.Username() + " - " + playerId + " to " + role)
}

func IsPlayerPrivileged(player proxy.Player) bool {
	role := GetPlayerRole(player)
	return role == "admin" || role == "builder" || role == "moderator"
}

func GetPlayerRank(player proxy.Player) string {
	playerId := player.ID().String()

	rankInterface := database.GetPlayerDataField(playerId, "rank")
	rank := "default"
	if rankInterface == nil {
		SetPlayerRank(player, "default")
	} else {
		rank = rankInterface.(string)
	}

	_, valid := validRanks[rank]
	if !valid {
		rank = "default"
		SetPlayerRank(player, "default")
	}

	return rank
}

func SetPlayerRank(player proxy.Player, rank string) {
	playerId := player.ID().String()

	_, valid := validRanks[rank]
	if !valid {
		logger.Error("Invalid rank: ", rank)
		return
	}

	database.SetPlayerDataField(playerId, "rank", rank)
	logger.Info("Changed permission rank for " + player.Username() + " - " + playerId + " to " + rank)
}
