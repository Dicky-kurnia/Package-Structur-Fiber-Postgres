package repository

import (
	"context"
	"go-fiber-postgres/model"
)

type AuthRepository interface {
	FetchSales(ctx context.Context, username string) (*model.Sales, error)
}
