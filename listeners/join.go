package listeners

import (
	"github.com/robinbraemer/event"
	"github.com/team-vesperis/vesperis-proxy/permission"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func registerJoinListener() {
	manager := p.Event()
	event.Subscribe(manager, 0, onLogin())
	event.Subscribe(manager, 0, onSpawn())

	logger.Info("Registered join listeners.")
}

func onLogin() func(*proxy.LoginEvent) {
	return func(event *proxy.LoginEvent) {

		// TODO: check if player is banned -> refuse login
	}
}

func onSpawn() func(*proxy.PostLoginEvent) {
	return func(event *proxy.PostLoginEvent) {
		player := event.Player()
		role := permission.GetPlayerRole(player.ID().String())
		rank := permission.GetPlayerRank(player.ID().String())

		player.SendMessage(&component.Text{
			Content: "Welcome to VesperisMC",
			S: component.Style{
				Color: color.Green,
			},
		})

		player.SendMessage(&component.Text{
			Content: "Your role: ",
			S: component.Style{
				Color: color.Gray,
			},
			Extra: []component.Component{
				&component.Text{
					Content: role,
					S: component.Style{
						Color: color.Aqua,
					},
				},
			},
		})

		player.SendMessage(&component.Text{
			Content: "Your rank: ",
			S: component.Style{
				Color: color.Gray,
			},
			Extra: []component.Component{
				&component.Text{
					Content: rank,
					S: component.Style{
						Color: color.Aqua,
					},
				},
			},
		})
	}
}
