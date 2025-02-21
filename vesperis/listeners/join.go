package listeners

import (
	"fmt"
	"time"

	"github.com/robinbraemer/event"
	"github.com/team-vesperis/vesperis-proxy/vesperis/ban"
	"github.com/team-vesperis/vesperis-proxy/vesperis/permission"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func registerJoinListeners() {
	event.Subscribe(manager, 0, onLogin())
	event.Subscribe(manager, 0, onPostLogin())

	logger.Info("Registered join listeners.")
}

func onLogin() func(*proxy.LoginEvent) {
	return func(event *proxy.LoginEvent) {
		player := event.Player()

		if ban.IsPlayerBanned(player) {
			reason := ban.GetBanReason(player)

			if ban.IsPlayerPermanentlyBanned(player) {
				event.Deny(&component.Text{
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

			} else {
				timeComponent := component.Text{}
				duration := time.Until(ban.GetBanExpiration(player))

				if duration.Seconds() < 1 {
					timeComponent = component.Text{
						Content: "\n\nYour ban has just expired. Please try again in a moment.",
						S: component.Style{
							Color: color.Aqua,
						},
					}

				} else {
					hours := int(duration.Hours())
					days := hours / 24
					hours = hours % 24
					minutes := int(duration.Minutes()) % 60
					seconds := int(duration.Seconds()) % 60

					timeComponent = component.Text{
						Content: "\n\nYou are still banned for " + fmt.Sprintf("%d days, %d hours, %d minutes and %d seconds", days, hours, minutes, seconds),
						S: component.Style{
							Color: color.Aqua,
						},
					}
				}

				event.Deny(&component.Text{
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
						&timeComponent,
					},
				})
			}
		}
	}
}

func onPostLogin() func(*proxy.PostLoginEvent) {
	return func(event *proxy.PostLoginEvent) {
		player := event.Player()
		role := permission.GetPlayerRole(player)
		rank := permission.GetPlayerRank(player)

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
