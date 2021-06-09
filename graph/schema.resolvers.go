package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ktp-fix/graph/generated"
	"ktp-fix/graph/model"
	"ktp-fix/internal/handlers"
	"ktp-fix/internal/models"
)

func (r *mutationResolver) CreateKtp(ctx context.Context, input models.NewKtp) (*models.Ktp, error) {
	return handlers.CreateKtpHandler(ctx, input)
}

func (r *mutationResolver) UpdateKtp(ctx context.Context, id string, input *models.NewKtp) (bool, error) {
	return handlers.UpdateKtp(ctx, id, input)
}

func (r *mutationResolver) DeleteKtp(ctx context.Context, id string) (bool, error) {
	return handlers.DeleteKtpHandler(ctx, id)
}

func (r *queryResolver) GetAllKtp(ctx context.Context) ([]*models.Ktp, error) {
	return handlers.GetAllKtp(ctx)
}

func (r *queryResolver) GetKtp(ctx context.Context, id string) (*models.Ktp, error) {
	return handlers.GetKtpHandler(ctx, id)
}

func (r *queryResolver) PaginateKtp(ctx context.Context, input model.Pagination) (*model.PaginationResultKtp, error) {
	return handlers.PaginateKtpHandler(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
