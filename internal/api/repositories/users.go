package repositories

import (
	"context"

	"gorm.io/gorm"

	"github.com/vanbien2402/first-web-demo/internal/api/domain"
	"github.com/vanbien2402/first-web-demo/internal/api/models"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindByUserName(ctx context.Context, userName string) (res *models.User, err error) {
	res = new(models.User)
	err = r.db.WithContext(ctx).Model(res).First(res, "user_name=?", userName).Error
	return
}

func (r *userRepository) FindByID(ctx context.Context, id string) (res *models.User, err error) {
	res = new(models.User)
	err = r.db.WithContext(ctx).Model(res).First(res, "id=?", id).Error
	return
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Model(user).
		Select([]string{"user_name", "password", "email", "version"}).
		Updates(user).
		Error
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
}

func (r *userRepository) CheckExist(ctx context.Context, userName string) bool {
	var count int64
	r.db.WithContext(ctx).Model(&models.User{}).Where("user_name = ?", userName).Count(&count)
	return count > 0
}

//NewUserRepository init user repository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
