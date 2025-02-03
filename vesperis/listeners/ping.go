package listeners

import (
	"strconv"

	"github.com/robinbraemer/event"
	"github.com/team-vesperis/vesperis-proxy/vesperis/utils"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/ping"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/favicon"
	"go.minekube.com/gate/pkg/util/uuid"
)

var fav favicon.Favicon

func registerPingListener() {
	event.Subscribe(manager, 0, onPing())
	loadFavicon()
}

func loadFavicon() {
	favi, err := favicon.FromFile("logo.png")
	if err != nil {
		logger.Error("Error loading logo. - ", err)
	}
	fav = favi
}

func onPing() func(*proxy.PingEvent) {
	return func(event *proxy.PingEvent) {

		ping := &ping.ServerPing{
			Version: ping.Version{
				Name:     "1.21.4",
				Protocol: 769,
			},

			Players: &ping.Players{
				Online: utils.GetPlayerCount(),
				Max:    200,
				Sample: []ping.SamplePlayer{
					{
						Name: "{\"text\":\"VesperisMC\",\"color\":\"#FFB108\"}",
						ID:   uuid.New(),
					},
					{
						Name: "",
						ID:   uuid.New(),
					},
					{
						Name: " - There are §b" + strconv.Itoa(utils.GetPlayerCount()) + "§f players online.",
						ID:   uuid.New(),
					},
					{
						Name: " - Vesperis Proxy: §b" + utils.GetVesperisProxyVersion(),
						ID:   uuid.New(),
					},
					{
						Name: " - Check our website§b www.vesperis.net §ffor the latest news & more!",
						ID:   uuid.New(),
					},
					{
						Name: "",
						ID:   uuid.New(),
					},
					{
						Name: "§6play.vesperis.net",
						ID:   uuid.New(),
					},
				},
			},

			Description: &component.Text{
				Content: "VesperisMC",
				S: component.Style{
					Color: utils.GetColorTitle(),
				},
				Extra: []component.Component{
					&component.Text{
						Content: "\n",
					},
					&component.Text{
						Content: "play.vesperis.net",
						S: component.Style{
							Color: utils.GetColorUnderTitle(),
						},
					},
				},
			},

			Favicon: fav,
		}

		event.SetPing(ping)
	}
}
