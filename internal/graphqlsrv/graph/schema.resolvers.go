package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/huyntsgs/graphql-project/internal/graphqlsrv/graph/generated"
	"github.com/huyntsgs/graphql-project/internal/graphqlsrv/graph/model"
)

func (r *queryResolver) Projects(ctx context.Context, last *int) (*model.Projects, error) {
	desc := "test"
	if last != nil {
		return &model.Projects{[]*model.Node{&model.Node{Name: "proj1", Description: &desc, ForksCount: 1},
			&model.Node{Name: "proj2", Description: &desc, ForksCount: 2},
			&model.Node{Name: "proj3", Description: &desc, ForksCount: 3},
			&model.Node{Name: "proj4", Description: &desc, ForksCount: 4},
			&model.Node{Name: "proj5", Description: &desc, ForksCount: 5},
			&model.Node{Name: "proj6", Description: &desc, ForksCount: 6}}}, nil
	} else {
		return &model.Projects{[]*model.Node{&model.Node{Name: "proj1", Description: &desc, ForksCount: 1},
			&model.Node{Name: "proj2", Description: &desc, ForksCount: 2},
			&model.Node{Name: "proj3", Description: &desc, ForksCount: 3},
			&model.Node{Name: "proj4", Description: &desc, ForksCount: 4},
			&model.Node{Name: "proj5", Description: &desc, ForksCount: 5}}}, nil
	}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
