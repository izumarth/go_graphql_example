package services

import (
	"context"

	"github.com/izumarth/go-graphql-example/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner string, name string) (*model.Repository, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}

type Services interface {
	UserService
	RepoService
	IssueService
}

type services struct {
	*userService
	*repoService
	*issueService
}

func New(
	exec boil.ContextExecutor,
) Services {
	return &services{
		userService:  &userService{exec: exec},
		repoService:  &repoService{exec: exec},
		issueService: &issueService{exec: exec},
	}
}
