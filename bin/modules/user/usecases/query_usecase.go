package usecases

import (
	"context"

	"github.com/Padliwinata/iam-mfa/bin/modules/user"
	"github.com/Padliwinata/iam-mfa/bin/modules/user/models"
	"github.com/Padliwinata/iam-mfa/bin/pkg/errors"
	"github.com/Padliwinata/iam-mfa/bin/pkg/utils"
)

type queryUsecase struct {
	userRepositoryQuery user.SqliteRepositoryQuery
}

func NewQueryUsecase(mq user.SqliteRepositoryQuery) user.UsecaseQuery {
	return &queryUsecase{
		userRepositoryQuery: mq,
	}
}

func (q queryUsecase) GetUser(ctx context.Context, userId string) utils.Result {
	var result utils.Result

	queryRes := <-q.userRepositoryQuery.FindOne(ctx, userId)

	if queryRes.Error != nil {
		errObj := errors.InternalServerError("Internal server error")
		result.Error = errObj
		return result
	}

	user := queryRes.Data.(models.User)
	res := models.GetUserResponse{
		Id:       user.ID.String(),
		Username: user.Name,
		Email:    user.Email,
	}
	result.Data = res
	return result
}
