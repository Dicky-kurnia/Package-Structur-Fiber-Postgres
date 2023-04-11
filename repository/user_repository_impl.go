package repository

import (
	"context"
	"errors"
	"go-fiber-postgres/model"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}

func (u userRepository) GetSalesIdByUsername(ctx context.Context, username string) (id uint, err error) {
	if err = u.DB.WithContext(ctx).Raw("SELECT id FROM sales WHERE username = ?", username).Scan(&id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return id, errors.New(model.NOT_FOUND_ERR_TYPE)
		}
		return id, err
	}
	return id, err
}

func (u userRepository) GetUserProfile(ctx context.Context, username string) (response *model.GetUserProfileResponse, err error) {
	if err = u.DB.WithContext(ctx).Raw("SELECT id as sales_id,username,name,role FROM sales WHERE username = ?", username).Scan(&response).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, errors.New(model.NOT_FOUND_ERR_TYPE)
		}
		return response, err
	}
	if response == nil {
		return nil, errors.New(model.SALES_NOT_FOUND)
	}
	return response, err
}

func (u userRepository) ChangePassword(ctx context.Context, username, password string) (string, error) {
	returning := ""
	execSQL := `UPDATE sales SET password = ? WHERE username = ? RETURNING username`
	err := u.DB.WithContext(ctx).Raw(execSQL, password, username).Scan(&returning).Error
	return returning, err
}
