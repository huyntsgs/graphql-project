package service

import (
	"github.com/huyntsgs/graphql-project/internal/repositories/gitlab"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/huyntsgs/graphql-project/internal/service/entities"
)

type IGitlabService interface {
	GetNodeInfo(n int) (*entities.NodeSummary, error)
}
type GitlabService struct {
	GitlabRepo gitlab.IGitlabRepo
	Logger     *zap.SugaredLogger
}

// using one param in new func helps reducing
// minimal change when adding more data field to service struct
// only needs to add to params struct and function signature does not change
type GitlabServiceParams struct {
	fx.In
	GitlabRepo gitlab.IGitlabRepo
	Logger     *zap.SugaredLogger
}

func NewGitlabService(params GitlabServiceParams) IGitlabService {
	return GitlabService{GitlabRepo: params.GitlabRepo, Logger: params.Logger}
}

func (service GitlabService) GetNodeInfo(n int) (*entities.NodeSummary, error) {

	response, err := service.GitlabRepo.GetLastProjectsInfo(n)
	if err != nil {
		service.Logger.Error("[GitlabService] - GetNodeInfo err: %v", err)
		return nil, err
	}

	nodeInfos := response.Projects["nodes"]

	nodeSummary := entities.NodeSummary{}
	if len(nodeInfos) == 0 {
		service.Logger.Info("[GitlabService] - GetNodeInfo empty data")
		return &nodeSummary, nil
	}

	for i, val := range nodeInfos {
		if i == len(nodeInfos)-1 {
			nodeSummary.Names = nodeSummary.Names + val.Name
		} else {
			nodeSummary.Names = nodeSummary.Names + val.Name + ","
		}

		nodeSummary.TotalForksCount += val.ForksCount
	}

	service.Logger.Info("[GitlabService] - GetNodeInfo node data: ", nodeSummary)
	return &nodeSummary, nil
}
