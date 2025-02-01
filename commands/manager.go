package commands

import (
	"strings"

	"github.com/team-vesperis/vesperis-proxy/permission"
	"github.com/team-vesperis/vesperis-proxy/vanish"
	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.uber.org/zap"
)

var p *proxy.Proxy
var logger *zap.SugaredLogger

func InitializeCommands(pr *proxy.Proxy, log *zap.SugaredLogger) {
	p = pr
	logger = log
	registerPermissionCommand()
	registerVanishCommand()
	registerBanCommand()
	registerTempBanCommand()
	registerUnBanCommand()
}

func requireAdmin() brigodier.RequireFn {
	return command.Requires(func(context *command.RequiresContext) bool {
		player := getPlayerFromSource(context.Source)

		if player != nil {
			return permission.GetPlayerRole(player) == "admin"
		}

		return false
	})
}

func requireAdminOrModerator() brigodier.RequireFn {
	return command.Requires(func(context *command.RequiresContext) bool {
		player := getPlayerFromSource(context.Source)

		if player != nil {
			permission := permission.GetPlayerRole(player)
			return permission == "admin" || permission == "moderator"
		}

		return false

	})
}

func requireStaff() brigodier.RequireFn {
	return command.Requires(func(context *command.RequiresContext) bool {
		player := getPlayerFromSource(context.Source)

		if player != nil {
			return permission.IsPlayerPrivileged(player)
		}

		return false
	})
}

func getPlayerTarget(playerName string, context command.Context) proxy.Player {
	player := p.PlayerByName(playerName)

	if player == nil {
		context.SendMessage(&component.Text{
			Content: "Player not found.",
			S: component.Style{
				Color: color.Red,
			},
		})

		return nil
	}

	return player
}

func getPlayerFromSource(source command.Source) proxy.Player {
	player, ok := source.(proxy.Player)

	if !ok {
		return nil
	}

	return player
}

func suggestPlayers() brigodier.SuggestionProvider {
	return command.SuggestFunc(func(context *command.Context, builder *brigodier.SuggestionsBuilder) *brigodier.Suggestions {
		remaining := builder.RemainingLowerCase

		players := make([]proxy.Player, 0)
		for _, player := range p.Players() {
			if sourcePlayer, ok := context.Source.(proxy.Player); ok {
				if permission.IsPlayerPrivileged(sourcePlayer) || vanish.IsPlayerVanished(player.ID().String()) {
					if strings.HasPrefix(strings.ToLower(player.Username()), remaining) {
						players = append(players, player)
					}
				}
			} else {
				if strings.HasPrefix(strings.ToLower(player.Username()), remaining) {
					players = append(players, player)
				}
			}
		}

		if len(players) != 0 {
			for _, player := range players {
				builder.Suggest(player.Username())
			}
		}

		return builder.Build()
	})
}
