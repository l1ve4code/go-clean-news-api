package user

import (
	"go-clean-news-api/internal/domain/entity"
	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *userStorage{
	return &userStorage{db: db}
}

func (bs *userStorage) GetOne(id int) *entity.User {
	var user *entity.User
	bs.db.Find(&user, id)
	return user
}
func (bs *userStorage) GetAll() []*entity.User {
	return nil
}
func (bs *userStorage) Create(u *entity.User) *entity.User {
	bs.db.Create(u)
	return u
}
func (bs *userStorage) Delete(u *entity.User) error {
	bs.db.Delete(u)
	return nil
}

func (bs *userStorage) FindByEmail(email string) *entity.User{
	var user entity.User
	bs.db.Where("email = ?", email).First(&user)
	return &user
}