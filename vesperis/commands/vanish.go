package commands

import (
	"github.com/team-vesperis/vesperis-proxy/vesperis/utils"
	"github.com/team-vesperis/vesperis-proxy/vesperis/vanish"
	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
)

func registerVanishCommand() {
	p.Command().Register(vanishCommand("vanish"))
	p.Command().Register(vanishCommand("v"))
	logger.Info("Registered vanish command.")
}

func vanishCommand(name string) brigodier.LiteralNodeBuilder {
	return brigodier.Literal(name).
		Then(brigodier.Literal("on").Executes(turnVanishOn())).
		Then(brigodier.Literal("off").Executes(turnVanishOff())).
		Executes(checkIfVanished()).
		Requires(requireStaff())
}

func checkIfVanished() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		player := getPlayerFromSource(context.Source)
		if player == nil {
			context.SendMessage(&component.Text{
				Content: "Only players can use the /vanish command.",
				S: component.Style{
					Color: color.Red,
				},
			})

			return nil
		}

		vanished := vanish.IsPlayerVanished(player)
		if vanished {
			player.SendMessage(&component.Text{
				Content: "You are vanished.",
				S: component.Style{
					Color: color.Aqua,
				},
			})
		} else {
			player.SendMessage(&component.Text{
				Content: "You are not vanished.",
				S: component.Style{
					Color: color.Aqua,
				},
			})
		}

		return nil
	})
}

func turnVanishOn() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		player := getPlayerFromSource(context.Source)
		if player == nil {
			context.SendMessage(&component.Text{
				Content: "Only players can use the /vanish command.",
				S: component.Style{
					Color: color.Red,
				},
			})

			return nil
		}

		if vanish.IsPlayerVanished(player) {
			player.SendMessage(&component.Text{
				Content: "You are already vanished.",
				S: component.Style{
					Color: utils.GetColorOrange(),
				},
			})

			return nil
		}

		vanish.SetPlayerVanished(player, true)
		player.SendMessage(&component.Text{
			Content: "You have vanished.",
			S: component.Style{
				Color: color.LightPurple,
			},
		})

		return nil
	})
}

func turnVanishOff() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		player := getPlayerFromSource(context.Source)
		if player == nil {
			context.SendMessage(&component.Text{
				Content: "Only players can use the /vanish command.",
				S: component.Style{
					Color: color.Red,
				},
			})

			return nil
		}

		if !vanish.IsPlayerVanished(player) {
			player.SendMessage(&component.Text{
				Content: "You are not vanished.",
				S: component.Style{
					Color: utils.GetColorOrange(),
				},
			})

			return nil
		}

		vanish.SetPlayerVanished(player, false)
		player.SendMessage(&component.Text{
			Content: "You have un-vanished.",
			S: component.Style{
				Color: color.LightPurple,
			},
		})

		return nil
	})
}
