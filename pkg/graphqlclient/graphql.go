package graphqlclient

import (
	"github.com/spf13/viper"

	"github.com/huyntsgs/graphql-project/config"
	"github.com/machinebox/graphql"
)

func NewGraphqlClient() *graphql.Client {
	grpClient := graphql.NewClient(viper.GetString(config.GITLAB_GRAPHQL_URL))
	return grpClient
}
