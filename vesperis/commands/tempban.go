package commands

import (
	"time"

	"github.com/team-vesperis/vesperis-proxy/vesperis/ban"
	"github.com/team-vesperis/vesperis-proxy/vesperis/permission"
	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
)

func registerTempBanCommand() {
	p.Command().Register(tempBanCommand())
	logger.Info("Registered temp ban command.")
}

func tempBanCommand() brigodier.LiteralNodeBuilder {
	return brigodier.Literal("tempban").
		Executes(incorrectTempBanCommandUsage()).
		Then(brigodier.Argument("player", brigodier.SingleWord).
			Executes(incorrectTempBanCommandUsage()).
			Suggests(suggestPlayers()).
			Then(brigodier.Argument("time_amount", brigodier.Int).
				Executes(incorrectTempBanCommandUsage()).
				Then(brigodier.Literal("seconds").
					Executes(incorrectTempBanCommandUsage()).
					Then(brigodier.Argument("reason", brigodier.StringPhrase).
						Executes(tempBanPlayer(time.Second)))).
				Then(brigodier.Literal("minutes").
					Executes(incorrectTempBanCommandUsage()).
					Then(brigodier.Argument("reason", brigodier.StringPhrase).
						Executes(tempBanPlayer(time.Minute)))).
				Then(brigodier.Literal("hours").
					Executes(incorrectTempBanCommandUsage()).
					Then(brigodier.Argument("reason", brigodier.StringPhrase).
						Executes(tempBanPlayer(time.Hour)))).
				Then(brigodier.Literal("days").
					Executes(incorrectTempBanCommandUsage()).
					Then(brigodier.Argument("reason", brigodier.StringPhrase).
						Executes(tempBanPlayer(time.Hour * 24)))))).
		Requires(requireAdminOrModerator())
}

func tempBanPlayer(time_type time.Duration) brigodier.Command {
	return command.Command(func(context *command.Context) error {
		playerName := context.String("player")
		player := getPlayerTarget(playerName, *context)
		if player == nil {
			return nil
		}

		if permission.GetPlayerRole(player) == "admin" {
			context.SendMessage(&component.Text{
				Content: "You are not allowed to ban admins.",
				S: component.Style{
					Color: color.Red,
				},
			})

			return nil
		}

		reason := context.String("reason")
		time_amount := context.Int("time_amount")
		ban.TempBanPlayer(player, reason, uint16(time_amount), time_type)

		return nil
	})
}

func incorrectTempBanCommandUsage() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		context.SendMessage(&component.Text{
			Content: "Incorrect usage: /tempban <player> <time_amount> <time_type> <reason>",
			S: component.Style{
				Color: color.Red,
			},
		})
		return nil
	})
}
