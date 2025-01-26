package commands

import (
	"github.com/team-vesperis/vesperis-proxy/permission"
	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
)

func registerPermissionCommand() {
	p.Command().Register(permissionCommand("permission"))
	p.Command().Register(permissionCommand("perm"))
	logger.Info("Registered permission command.")
}

func permissionCommand(name string) brigodier.LiteralNodeBuilder {
	return brigodier.Literal(name).
		Then(brigodier.Literal("set").
			Then(brigodier.Literal("role").
				Then(brigodier.Argument("player", brigodier.SingleWord).
					Then(brigodier.Literal("admin").
						Executes(setRole("admin"))).
					Then(brigodier.Literal("builder").
						Executes(setRole("builder"))).
					Then(brigodier.Literal("default").
						Executes(setRole("default"))).
					Then(brigodier.Literal("moderator").
						Executes(setRole("moderator"))).
					Executes(incorrectSetUsage()).
					Suggests(suggestPlayers())).
				Executes(incorrectSetUsage())).
			Executes(incorrectSetUsage()).
			Then(brigodier.Literal("rank").
				Then(brigodier.Argument("player", brigodier.SingleWord).
					Then(brigodier.Literal("champion").
						Executes(setRank("champion"))).
					Then(brigodier.Literal("default").
						Executes(setRank("default"))).
					Then(brigodier.Literal("elite").
						Executes(setRank("elite"))).
					Then(brigodier.Literal("legend").
						Executes(setRank("legend"))).
					Executes(incorrectSetUsage()).
					Suggests(suggestPlayers())).
				Executes(incorrectSetUsage())).
			Requires(requireAdmin())).
		Then(brigodier.Literal("get").
			Then(brigodier.Literal("rank").
				Then(brigodier.Argument("player", brigodier.SingleWord).
					Executes(getRank()).
					Suggests(suggestPlayers()))).
			Then(brigodier.Literal("role").
				Then(brigodier.Argument("player", brigodier.SingleWord).
					Executes(getRole()).
					Suggests(suggestPlayers()))).
			Executes(incorrectGetUsage())).
		Executes(incorrectFullUsage()).
		Requires(requireAdminOrModerator())
}

func setRole(role string) brigodier.Command {
	return command.Command(func(context *command.Context) error {
		playerName := context.String("player")
		player := getPlayerTarget(playerName, *context)
		if player == nil {
			return nil
		}

		permission.SetPlayerRole(player.ID().String(), role)
		context.SendMessage(&component.Text{
			Content: "Set role for player ",
			S:       component.Style{Color: color.Green},
			Extra: []component.Component{
				&component.Text{
					Content: playerName,
					S:       component.Style{Color: color.Aqua},
				},
				&component.Text{
					Content: " to ",
					S:       component.Style{Color: color.Green},
				},
				&component.Text{
					Content: role,
					S:       component.Style{Color: color.Gold},
				},
			},
		})

		return nil
	})
}

func setRank(rank string) brigodier.Command {
	return command.Command(func(context *command.Context) error {
		playerName := context.String("player")
		player := getPlayerTarget(playerName, *context)
		if player == nil {
			return nil
		}

		permission.SetPlayerRank(player.ID().String(), rank)
		context.SendMessage(&component.Text{
			Content: "Set rank for player ",
			S:       component.Style{Color: color.Green},
			Extra: []component.Component{
				&component.Text{
					Content: playerName,
					S:       component.Style{Color: color.Aqua},
				},
				&component.Text{
					Content: " to ",
					S:       component.Style{Color: color.Green},
				},
				&component.Text{
					Content: rank,
					S:       component.Style{Color: color.Gold},
				},
			},
		})

		return nil
	})
}

func getRank() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		playerName := context.String("player")
		player := getPlayerTarget(playerName, *context)
		if player == nil {
			return nil
		}

		rank := permission.GetPlayerRank(player.ID().String())
		context.SendMessage(&component.Text{
			Content: "The rank of ",
			S:       component.Style{Color: color.Green},
			Extra: []component.Component{
				&component.Text{
					Content: playerName,
					S:       component.Style{Color: color.Aqua},
				},
				&component.Text{
					Content: " is: ",
					S:       component.Style{Color: color.Green},
				},
				&component.Text{
					Content: rank,
					S:       component.Style{Color: color.Gold},
				},
			},
		})

		return nil
	})
}

func getRole() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		playerName := context.String("player")
		player := getPlayerTarget(playerName, *context)
		if player == nil {
			return nil
		}

		role := permission.GetPlayerRole(player.ID().String())
		context.SendMessage(&component.Text{
			Content: "The role of ",
			S:       component.Style{Color: color.Green},
			Extra: []component.Component{
				&component.Text{
					Content: playerName,
					S:       component.Style{Color: color.Aqua},
				},
				&component.Text{
					Content: " is: ",
					S:       component.Style{Color: color.Green},
				},
				&component.Text{
					Content: role,
					S:       component.Style{Color: color.Gold},
				},
			},
		})

		return nil
	})
}

func incorrectFullUsage() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		context.SendMessage(&component.Text{
			Content: "Incorrect usage:\n 1. /permission set role <player> <role>\n 2. /permission set rank <player> <rank>\n 3. /permission get role <player>\n 4. /permission get rank <player>",
			S:       component.Style{Color: color.Red},
		})
		return nil
	})
}

func incorrectSetUsage() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		context.SendMessage(&component.Text{
			Content: "Incorrect usage: /permission set role/rank <player> <role/rank>",
			S:       component.Style{Color: color.Red},
		})
		return nil
	})
}

func incorrectGetUsage() brigodier.Command {
	return command.Command(func(context *command.Context) error {
		context.SendMessage(&component.Text{
			Content: "Incorrect usage: /permission get <player>",
			S:       component.Style{Color: color.Red},
		})
		return nil
	})
}
