package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/team-vesperis/vesperis-proxy/vesperis/ban"
	"github.com/team-vesperis/vesperis-proxy/vesperis/commands"
	"github.com/team-vesperis/vesperis-proxy/vesperis/config"
	"github.com/team-vesperis/vesperis-proxy/vesperis/database"
	"github.com/team-vesperis/vesperis-proxy/vesperis/listeners"
	"github.com/team-vesperis/vesperis-proxy/vesperis/logger"
	"github.com/team-vesperis/vesperis-proxy/vesperis/permission"
	"github.com/team-vesperis/vesperis-proxy/vesperis/utils"

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
	logger := logger.GetLogger()
	logger.Info("Starting Vesperis Proxy...")

	config.InitializeConfig(logger)

	proxy.Plugins = append(proxy.Plugins,
		proxy.Plugin{
			Name: "VesperisProxy",
			Init: func(ctx context.Context, proxy *proxy.Proxy) error {
				return newVesperisProxy(proxy, logger).init()
			},
		},
	)

	handleShutdown(logger)
	gate.Execute()

	logger.Info("Successfully started Vesperis Proxy.")
}

func handleShutdown(logger *zap.SugaredLogger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		logger.Info("Stopping Vesperis Proxy...")

		close(ban.Quit)
		database.CloseDatabase()

		logger.Info("Successfully stopped Vesperis Proxy.")
		os.Exit(0)
	}()
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
	ban.InitializeBanManager(vp.logger)

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
