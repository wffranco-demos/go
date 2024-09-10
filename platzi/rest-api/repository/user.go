package repository

import (
	"context"

	"github.com/wffranco-demos/go/platzi/rest-api/models"
)

type UserRepository = Repository[models.User]

var implementation UserRepository

func SetUserRepository(repo UserRepository) {
	implementation = repo
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.Insert(ctx, user)
}

func GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetByID(ctx, id)
}

func CloseUserRepository() error {
	return implementation.Close()
}
