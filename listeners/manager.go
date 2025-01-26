package listeners

import (
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.uber.org/zap"
)

var p *proxy.Proxy
var logger *zap.SugaredLogger

func InitializeListeners(pr *proxy.Proxy, log *zap.SugaredLogger) {
	p = pr
	logger = log
	registerJoinListener()
}
