package repository

import (
	"context"
	"errors"
	"go-fiber-postgres/exception"
	"go-fiber-postgres/model"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{DB: db}
}

func (repo *authRepository) FetchSales(ctx context.Context, username string) (sales *model.Sales, err error) {
	if err = repo.DB.WithContext(ctx).
		Raw("SELECT * FROM sales WHERE username = ?", username).Scan(&sales).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sales, exception.DatabaseError{Message: "Record not found"}
		}
		return sales, err
	}
	if sales == nil {
		return sales, exception.DatabaseError{Message: "Record not found"}
	}
	return sales, nil
}
