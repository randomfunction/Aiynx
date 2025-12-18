package service

import (
	"context"
	"errors"
	"time"
	"user-api/db/sqlc"
	"user-api/internal/models"
	"user-api/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo      repository.UserRepository
	validator *validator.Validate
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	if dob.After(time.Now()) {
		return nil, errors.New("date of birth cannot be in the future")
	}

	arg := sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  pgtype.Date{Time: dob, Valid: true},
	}

	user, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return mapUserToResponse(user), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return mapUserToResponse(user), nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, u := range users {
		responses = append(responses, *mapUserToResponse(u))
	}
	return responses, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, req models.UpdateUserRequest) (*models.UserResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	if dob.After(time.Now()) {
		return nil, errors.New("date of birth cannot be in the future")
	}

	arg := sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: req.Name,
		Dob:  pgtype.Date{Time: dob, Valid: true},
	}

	user, err := s.repo.UpdateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return mapUserToResponse(user), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, int32(id))
}

func mapUserToResponse(user sqlc.User) *models.UserResponse {
	dob := user.Dob.Time
	age := CalculateAge(dob, time.Now())

	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Dob:       dob.Format("2006-01-02"),
		Age:       age,
		CreatedAt: user.CreatedAt.Time,
	}
}

func CalculateAge(dob time.Time, now time.Time) int {
	age := now.Year() - dob.Year()
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	return age
}
