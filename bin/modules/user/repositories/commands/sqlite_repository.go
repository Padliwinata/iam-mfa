package commands

import (
	"context"

	"gorm.io/gorm"

	"github.com/Padliwinata/iam-mfa/bin/modules/user/models"
	"github.com/Padliwinata/iam-mfa/bin/pkg/utils"
)

type querySQLiteRepository struct {
	DB *gorm.DB
}

func (q querySQLiteRepository) InsertOne(ctx context.Context, data models.User) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := q.DB.Create(&data)
		if result.Error != nil {
			output <- utils.Result{
				Error: result.Error,
			}
		}

		output <- utils.Result{
			Data: data,
		}
	}()

	return output
}
