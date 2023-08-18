package user

import (
	"context"

	"github.com/Padliwinata/iam-mfa/bin/modules/user/models"
	"github.com/Padliwinata/iam-mfa/bin/pkg/utils"
)

type UsecaseQuery interface {
	GetUser(ctx context.Context, userId string) utils.Result
}
type UsecaseCommand interface {
}

type SqliteRepositoryQuery interface {
	FindOne(ctx context.Context, userID string) <-chan utils.Result
}

type SqliteRepositoryCommand interface {
	InsertOne(ctx context.Context, data models.User) <-chan utils.Result
}
