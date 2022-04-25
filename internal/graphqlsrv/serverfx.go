package graphqlsrv

import (
	"context"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/huyntsgs/graphql-project/internal/graphqlsrv/graph"
	"github.com/huyntsgs/graphql-project/internal/graphqlsrv/graph/generated"
)

const defaultPort = "8080"

var Runner = fx.Invoke(
	startServer,
)

func startServer(lifecycle fx.Lifecycle, logger *zap.SugaredLogger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					port := viper.GetString("PORT")
					if port == "" {
						port = defaultPort
					}

					srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

					http.Handle("/", playground.Handler("GraphQL playground", "/query"))
					http.Handle("/query", srv)

					logger.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
					logger.Fatal(http.ListenAndServe(":"+port, nil))

				}()
				return nil
			},

			OnStop: func(context.Context) error {
				logger.Info("Disconected http://localhost:%s/ GraphQL playground", viper.GetString("PORT"))
				return nil
			},
		},
	)

}
