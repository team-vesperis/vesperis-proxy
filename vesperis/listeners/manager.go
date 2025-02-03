package listeners

import (
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/key"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/edition/java/proxy/message"
	"go.uber.org/zap"
)

var p *proxy.Proxy
var logger *zap.SugaredLogger
var manager event.Manager

func InitializeListeners(pr *proxy.Proxy, log *zap.SugaredLogger) {
	p = pr
	logger = log
	manager = p.Event()
	registerJoinListeners()
	registerPingListener()

	keyt, _ := key.Make("minecraft", "brand")
	logger.Info(message.MinecraftChannelIdentifier{
		Key: key.New(key.MinecraftNamespace, "brand"),
	},
	)
	logger.Info(keyt.String())
	event.Subscribe(manager, 0, onPluginMessage())
}

func onPluginMessage() func(*proxy.PluginMessageEvent) {
	return func(event *proxy.PluginMessageEvent) {
		logger.Info(event.Data())
		if (event.Identifier() == &message.MinecraftChannelIdentifier{Key: key.New(key.MinecraftNamespace, "brand")}) {
			logger.Info(event.Data(), event.Target(), event.Source(), event.Identifier())
		}
	}
}
