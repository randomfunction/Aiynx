package repository

import (
	"context"
	"user-api/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, id int32) (sqlc.User, error)
	ListUsers(ctx context.Context) ([]sqlc.User, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type sqlRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) UserRepository {
	return &sqlRepository{queries: q}
}

func (r *sqlRepository) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *sqlRepository) GetUserByID(ctx context.Context, id int32) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *sqlRepository) ListUsers(ctx context.Context) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx)
}

func (r *sqlRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *sqlRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
