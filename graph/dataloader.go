package graph

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/izumarth/go-graphql-example/graph/model"
	"github.com/izumarth/go-graphql-example/graph/services"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

type userBatcher struct {
	Srv services.Services
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}
	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

func (u *userBatcher) BatchGetUsers(
	ctx context.Context,
	IDs []string,
) []*dataloader.Result[*model.User] {

	// 初期化
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	indexs := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	users, err := u.Srv.ListUsersByID(ctx, IDs)
	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if err != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexs[user.ID]] = rsl
	}
	return results
}
