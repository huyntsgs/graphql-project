package handlers

import (
	"context"

	"github.com/huyntsgs/graphql-project/internal/service"
	"github.com/huyntsgs/graphql-project/pkg/commandsrv"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type GitlabHandler struct {
	Serivce service.IGitlabService
	Logger  *zap.SugaredLogger
}

type GitlabHandlerParams struct {
	fx.In
	Serivce service.IGitlabService
	// other services below if any
	Logger *zap.SugaredLogger
}

func NewGitlabHandler(params GitlabHandlerParams) GitlabHandler {
	return GitlabHandler{Serivce: params.Serivce, Logger: params.Logger}
}

func (h GitlabHandler) GetLastProjects() commandsrv.CommandFunc {
	return func(ctx context.Context) {
		n := ctx.Value("n").(int)
		node, _ := h.Serivce.GetNodeInfo(n)
		h.Logger.Info("[handlers] - GitlabHandler - node data ", node)
	}
}
