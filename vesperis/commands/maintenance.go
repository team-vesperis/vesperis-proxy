package commands

import (
	"go.minekube.com/brigodier"
)

func registerMaintenanceCommand() {
	p.Command().Register(maintenanceCommand("maintenance"))
	p.Command().Register(maintenanceCommand("mnt"))
	logger.Info("Registered maintenance command.")
}

func maintenanceCommand(name string) brigodier.LiteralNodeBuilder {
	return brigodier.Literal(name).
		Requires(requireStaff()).
		Then(brigodier.Literal("all").
			Then(brigodier.Literal("on")).
			Then(brigodier.Literal("off"))).
		Then(brigodier.Literal("server").
			Then(brigodier.Argument("server", brigodier.SingleWord)))
}
