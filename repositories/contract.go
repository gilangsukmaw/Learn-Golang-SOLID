package repositories

import (
	"context"
	"solid/models"
)

type DbOperations interface {
	Read(ctx context.Context) (models.User, error)
	Update(ctx context.Context, payload models.User, whereId uint64) bool
	Create(ctx context.Context, payload models.User) bool
	Delete(ctx context.Context, whereId uint64) bool
}
