package resolver

import (
	"context"
	"fmt"

	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/model"
)

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByIDRequest) (*model.GetUserByIDResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *resolver }
