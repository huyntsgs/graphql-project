package controller

import (
	"github.com/huyntsgs/graphql-project/internal/controller/handlers"
	"github.com/huyntsgs/graphql-project/pkg/commandsrv"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	setupRouter,
	handlers.NewGitlabHandler,
)

type RoutingParams struct {
	fx.In
	GitlabHandler handlers.GitlabHandler
	// other handlers below if any
}

type RouteFuncGroup struct {
	fx.Out
	RouteFunc *commandsrv.RouteFunc
}

func setupRouter(params RoutingParams) RouteFuncGroup {

	routeFunc := map[string]commandsrv.CommandInfo{
		"get_lastproject": commandsrv.CommandInfo{CommandId: "get_lastproject",
			CommandFunc: params.GitlabHandler.GetLastProjects()},
	}

	return RouteFuncGroup{RouteFunc: &commandsrv.RouteFunc{routeFunc}}
}
