package services

import (
	"context"

	"github.com/saki-engineering/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner string, name string) (*model.Repository, error)
}

type Services interface {
	UserService
	RepoService
}

type services struct {
	*userService
	*repoService
}

func New(
	exec boil.ContextExecutor,
) Services {
	return &services{
		userService: &userService{exec: exec},
		repoService: &repoService{exec: exec},
	}
}
