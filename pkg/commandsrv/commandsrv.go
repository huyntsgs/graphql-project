package commandsrv

import (
	"context"

	"go.uber.org/fx"
	//"go.uber.org/fx"
)

var Module = fx.Provide(
	NewCommandServer,
)

type RouteFunc struct {
	CommandInfos map[string]CommandInfo
}

type CommandFunc func(context.Context)

type CommandInfo struct {
	CommandId   string
	CommandName string
	CommandData map[string]interface{}
	CommandFunc CommandFunc
}

type CommandServer struct {
	RouteFn RouteFunc
}

func NewCommandServer(RouteFn RouteFunc) CommandServer {
	return CommandServer{RouteFn: RouteFn}
}

func (command CommandServer) Route(ctx context.Context, cmd CommandInfo) {
	if len(command.RouteFn.CommandInfos) == 0 {
		return
	}

	for _, cmdInfo := range command.RouteFn.CommandInfos {
		if cmdInfo.CommandId == cmd.CommandId {
			ctx.Value(cmd.CommandData)
			cmdInfo.CommandFunc(ctx)
		}
	}
}
