package queries

import (
	"context"

	"gorm.io/gorm"

	"github.com/Padliwinata/iam-mfa/bin/modules/user/models"
	"github.com/Padliwinata/iam-mfa/bin/pkg/utils"
)

type querySQLiteRepository struct {
	DB *gorm.DB
}

func (q querySQLiteRepository) FindOne(ctx context.Context, userID string) <-chan utils.Result {
	output := make(chan utils.Result)
	var user models.User

	go func() {
		defer close(output)

		result := q.DB.First(&user, "id = ?", userID)
		if result.Error != nil {
			output <- utils.Result{
				Error: result.Error,
			}
		}

		output <- utils.Result{
			Data: user,
		}
	}()

	return output
}
