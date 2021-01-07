package resolver

import (
	"context"

	"github.com/takuya911/project-services/services/gql/graph/generated"
	"github.com/takuya911/project-services/services/gql/graph/model"
	definition "github.com/takuya911/project-user-definition"
)

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByIDRequest) (*model.GetUserByIDResponse, error) {
	res, err := r.userClient.GetByID(ctx, &definition.GetUserRequest{
		Uuid: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &model.GetUserByIDResponse{User: &model.User{
		ID:              res.GetUuid(),
		Name:            res.GetName(),
		Email:           res.GetEmail(),
		Password:        res.GetPassword(),
		TelephoneNumber: res.GetTelephoneNumber(),
		Gender:          res.GetGender(),
		CreatedAt:       res.GetCreatedAt(),
		UpdatedAt:       res.GetUpdatedAt(),
	}}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *resolver }