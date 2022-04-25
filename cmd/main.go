package main

import (
	"flag"
	"log"

	"github.com/huyntsgs/graphql-project/internal/controller"

	"github.com/huyntsgs/graphql-project/pkg/commandsrv"
	"github.com/huyntsgs/graphql-project/pkg/logger"

	"go.uber.org/fx"

	"github.com/huyntsgs/graphql-project/internal/repositories/gitlab"
	"github.com/huyntsgs/graphql-project/internal/service"
	"github.com/huyntsgs/graphql-project/pkg/graphqlclient"

	"github.com/huyntsgs/graphql-project/config"
	"github.com/huyntsgs/graphql-project/internal/graphqlsrv"
)

var cmdId string
var configPath string

func init() {
	flag.StringVar(&cmdId, "cmdid", config.CMD_GRAPHQL_SERVER, "Command id to execute")
	flag.StringVar(&configPath, "config_path", config.DEFAULT_CONFIG_FILE, "Config file path")
}

func main() {

	flag.Parse()

	log.Println("configPath", configPath)
	config.ParseConfig(configPath)
	var fxApp *fx.App

	if cmdId == "GRAPHQL_SERVER" {
		fxApp = fx.New(
			logger.Module,
			graphqlsrv.Runner,
		)
		if fxApp.Err() != nil {
			panic(fxApp.Err())
		}

		fxApp.Run()
	}

	if cmdId == "GET_PROJECTS" {
		fxApp = fx.New(
			logger.Module,
			graphqlclient.Module,
			gitlab.Module,
			service.Module,
			controller.Module,
			commandsrv.Module,
			fx.Invoke(func(srv service.IGitlabService) {
				srv.GetNodeInfo(5)
			},
			),
		)
		if fxApp.Err() != nil {
			panic(fxApp.Err())
		}
	}

}
