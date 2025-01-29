package commands

import (
	"github.com/team-vesperis/vesperis-proxy/ban"
	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
)

func registerUnBanCommand() {
	p.Command().Register(unBanCommand())
	logger.Info("Registered un-ban command.")
}

func unBanCommand() brigodier.LiteralNodeBuilder {
	return brigodier.Literal("unban").
		Then(brigodier.Argument("player", brigodier.SingleWord).
			Executes(unBanPlayer()).
			Suggests(suggestBannedPlayers())).
		Executes(incorrectUnBanCommandUsage()).
		Requires(requireAdminOrModerator())
}

func unBanPlayer() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		playerName := context.String("player")
		playerId := ban.GetBannedPlayerIdByName(playerName)
		if playerId == "" {
			context.SendMessage(&component.Text{
				Content: "Player not found.",
				S:       component.Style{Color: color.Red},
			})
			return nil
		}

		ban.UnBanPlayer(playerId)
		context.SendMessage(&component.Text{
			Content: "Player " + playerName + " has been unbanned.",
			S:       component.Style{Color: color.Green},
		})

		return nil
	})
}

func suggestBannedPlayers() brigodier.SuggestionProvider {
	return command.SuggestFunc(func(context *command.Context, builder *brigodier.SuggestionsBuilder) *brigodier.Suggestions {
		var list []string = ban.GetBannedPlayerNameList()
		for _, playerName := range list {
			builder.Suggest(playerName)
		}
		return builder.Build()
	})
}

func incorrectUnBanCommandUsage() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		context.SendMessage(&component.Text{
			Content: "Incorrect usage: /unban <player>",
			S: component.Style{
				Color: color.Red,
			},
		})
		return nil
	})
}
