package repositories

import (
	"context"
	"gorm.io/gorm"
	"solid/models"
)

type dbOperations struct {
	db *gorm.DB
}

func NewDbOperations(db *gorm.DB) DbOperations {
	return &dbOperations{
		db: db,
	}
}

func (r *dbOperations) Read(ctx context.Context) (models.User, error) {
	var user models.User

	r.db.Find(&user)

	return user, nil
}

func (r *dbOperations) Update(ctx context.Context, payload models.User, id uint64) bool {
	result := r.db.Model(&payload).Where("id = ?", id).Update("name", payload.Name).WithContext(ctx)

	if result.Error != nil {
		return false
	}

	return true
}

func (r *dbOperations) Create(ctx context.Context, payload models.User) bool {
	result := r.db.Create(&payload).WithContext(ctx) // pass pointer of data to Create

	if result.Error != nil {
		return false
	}
	return true
}

func (r *dbOperations) Delete(ctx context.Context, id uint64) bool {
	result := r.db.Where("id = ?", id).Delete(&models.User{}).WithContext(ctx)
	if result.Error != nil {
		return false
	}
	return true
}
