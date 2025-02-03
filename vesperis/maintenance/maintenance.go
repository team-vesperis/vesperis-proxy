package maintenance

import (
	"github.com/team-vesperis/vesperis-proxy/vesperis/permission"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func MaintenanceServer(server proxy.RegisteredServer) {
	server.Players().Range(func(player proxy.Player) bool {
		if !permission.IsPlayerPrivileged(player) {
			player.Disconnect(&component.Text{
				Content: "This server has gone into maintenance.",
			})
		}

		return true
	})
}
