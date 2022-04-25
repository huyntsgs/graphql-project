package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/huyntsgs/graphql-project/internal/repositories/gitlab"
	"github.com/huyntsgs/graphql-project/internal/repositories/models"
	"go.uber.org/zap"
)

func Test_GetNodeInfo_Then_Success(t *testing.T) {

	logger, _ := zap.NewProduction()
	mockGitlabRepo := gitlab.NewMockIGitlabRepo(t)

	projects := map[string][]models.NodeInfo{
		"nodes": []models.NodeInfo{
			models.NodeInfo{Name: "project1", ForksCount: 1},
			models.NodeInfo{Name: "project2", ForksCount: 1},
			models.NodeInfo{Name: "project3", ForksCount: 1},
		},
	}
	resp := &models.GetNodeResponse{Projects: projects}
	mockGitlabRepo.On("GetLastProjectsInfo", mock.Anything).Return(resp, nil)

	params := GitlabServiceParams{
		GitlabRepo: mockGitlabRepo,
		Logger:     logger.Sugar(),
	}

	service := NewGitlabService(params)
	res, _ := service.GetNodeInfo(5)

	assert.Equal(t, res.Names, "project1,project2,project3")
	assert.Equal(t, res.TotalForksCount, 3)

}

func Test_GetNodeInfo_Then_Error(t *testing.T) {

	logger, _ := zap.NewProduction()
	mockGitlabRepo := gitlab.NewMockIGitlabRepo(t)

	mockGitlabRepo.On("GetLastProjectsInfo", mock.Anything).Return(nil, errors.New("anyerror"))

	params := GitlabServiceParams{
		GitlabRepo: mockGitlabRepo,
		Logger:     logger.Sugar(),
	}

	service := NewGitlabService(params)
	_, err := service.GetNodeInfo(5)

	assert.NotNil(t, err)

}
