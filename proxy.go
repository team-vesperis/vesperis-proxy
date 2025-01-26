package main

import (
	"context"
	"fmt"
	"os"

	"github.com/team-vesperis/vesperis-proxy/commands"
	"github.com/team-vesperis/vesperis-proxy/config"
	"github.com/team-vesperis/vesperis-proxy/database"
	"github.com/team-vesperis/vesperis-proxy/listeners"
	"github.com/team-vesperis/vesperis-proxy/logger"
	"github.com/team-vesperis/vesperis-proxy/permission"
	"github.com/team-vesperis/vesperis-proxy/utils"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.uber.org/zap"
)

type VesperisProxy struct {
	*proxy.Proxy
	logger *zap.SugaredLogger
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error occurred:", r)
			fmt.Println("Press 'Enter' to exit...")
			fmt.Scanln()
			os.Exit(1)
		}
	}()

	logger.CreateLogger()
	log := logger.GetLogger()
	log.Info("Starting Vesperis Proxy...")

	config.InitializeConfig(log)

	proxy.Plugins = append(proxy.Plugins,
		proxy.Plugin{
			Name: "VesperisProxy",
			Init: func(ctx context.Context, proxy *proxy.Proxy) error {
				return newVesperisProxy(proxy, log).init()
			},
		},
	)

	gate.Execute()

	log.Info("Successfully started Vesperis Proxy.")
	fmt.Println("Press 'Enter' to exit...")
	fmt.Scanln()
}

func newVesperisProxy(proxy *proxy.Proxy, logger *zap.SugaredLogger) *VesperisProxy {
	return &VesperisProxy{
		Proxy:  proxy,
		logger: logger,
	}
}

func (vp *VesperisProxy) init() error {
	database.InitializeDatabase(vp.logger)

	permission.InitializePermissionManager(vp.logger)
	utils.InitializeUtils(vp.Proxy, vp.logger)

	vp.registerCommands()
	vp.registerEvents()

	return nil
}

func (vp *VesperisProxy) registerCommands() {
	vp.logger.Info("Registering commands...")
	commands.InitializeCommands(vp.Proxy, vp.logger)
	vp.logger.Info("Successfully registered commands.")
}

func (vp *VesperisProxy) registerEvents() {
	vp.logger.Info("Registering events...")
	listeners.InitializeListeners(vp.Proxy, vp.logger)
	vp.logger.Info("Successfully registered events.")
}
