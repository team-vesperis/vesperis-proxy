package ban

import (
	"fmt"
	"time"

	"github.com/team-vesperis/vesperis-proxy/database"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var Quit chan struct{} = make(chan struct{})

func InitializeBanManager(log *zap.SugaredLogger) {
	logger = log
	logger.Info("Initialized Ban Manager.")

	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				database.CheckTempBans()
			case <-Quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func IsPlayerBanned(player proxy.Player) bool {
	return database.IsPlayerBanned(player.ID().String())
}

func IsPlayerPermanentlyBanned(player proxy.Player) bool {
	return database.IsPlayerPermanentlyBanned(player.ID().String())
}

func GetBanReason(player proxy.Player) string {
	reason := database.GetBanReason(player.ID().String())
	if reason == "" {
		return "No reason provided."
	}

	return reason
}

func BanPlayer(player proxy.Player, reason string) {
	player.Disconnect(&component.Text{
		Content: "You are permanently banned from VesperisMC.",
		S: component.Style{
			Color: color.Red,
		},
		Extra: []component.Component{
			&component.Text{
				Content: "\n\nReason: " + reason,
				S: component.Style{
					Color: color.Gray,
				},
			},
		},
	})

	logger.Info("Player " + player.Username() + " - " + player.ID().String() + " has been banned. Reason: " + reason)
	database.BanPlayer(player.ID().String(), player.Username(), reason)
}

func TempBanPlayer(player proxy.Player, reason string, durationLength uint16, durationType time.Duration) {
	duration := time.Duration(durationLength) * durationType
	hours := int(duration.Hours())
	days := hours / 24
	hours = hours % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	time := fmt.Sprintf("%d days, %d hours, %d minutes and %d seconds", days, hours, minutes, seconds)

	player.Disconnect(&component.Text{
		Content: "You are temporarily banned from VesperisMC",
		S: component.Style{
			Color: color.Red,
		},
		Extra: []component.Component{
			&component.Text{
				Content: "\n\nReason: " + reason,
				S: component.Style{
					Color: color.Gray,
				},
			},
			&component.Text{
				Content: "\n\nYou are banned for " + time,
				S: component.Style{
					Color: color.Aqua,
				},
			},
		},
	})

	logger.Info("Player " + player.Username() + " - " + player.ID().String() + " has been temporarily banned. Reason: " + reason + " Time: " + time)
	database.TempBanPlayer(player.ID().String(), player.Username(), reason, durationLength, durationType)
}

func UnBanPlayer(playerId string) {
	logger.Info("Player with ID: " + playerId + " has been manually unbanned.")
	database.UnBanPlayer(playerId)
}

func GetBanExpiration(player proxy.Player) time.Time {
	return database.GetBanExpiration(player.ID().String())
}

func GetBannedPlayerNameList() []string {
	return database.GetBannedPlayerNameList()
}

func GetBannedPlayerIdByName(playerName string) string {
	return database.GetBannedPlayerIdByName(playerName)
}
