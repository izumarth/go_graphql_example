package services

import (
	"context"

	"github.com/izumarth/go-graphql-example/graph/db"
	"github.com/izumarth/go-graphql-example/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repoService struct {
	exec boil.ContextExecutor
}

func (r *repoService) GetRepoByFullName(
	ctx context.Context,
	owner string,
	name string,
) (*model.Repository, error) {
	repo, error := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)

	if error != nil {
		return nil, error
	}

	return convertRepository(repo), nil
}

func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Owner:     &model.User{ID: repo.Owner},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}
}
