package graphqlclient

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewGraphqlClient,
)
