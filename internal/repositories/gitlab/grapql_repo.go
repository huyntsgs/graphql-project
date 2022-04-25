package gitlab

import (
	"context"

	"github.com/huyntsgs/graphql-project/config"

	"go.uber.org/fx"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	"github.com/huyntsgs/graphql-project/internal/repositories/models"
	"github.com/machinebox/graphql"
)

type IGitlabRepo interface {
	GetLastProjectsInfo(n int) (*models.GetNodeResponse, error)
}

type GitlabRepo struct {
	Client *graphql.Client
	Logger *zap.SugaredLogger
}
type GitlabRepoParams struct {
	fx.In
	Client *graphql.Client
	Logger *zap.SugaredLogger
}

func NewGitlabRepo(params GitlabRepoParams) IGitlabRepo {
	return GitlabRepo{Client: params.Client, Logger: params.Logger}
}

func (repo GitlabRepo) GetLastProjectsInfo(n int) (*models.GetNodeResponse, error) {

	var response models.GetNodeResponse
	var request = graphql.NewRequest(viper.GetString(config.PROJECT_QUERY))
	request.Var("n", n)
	err := repo.Client.Run(context.Background(), request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
