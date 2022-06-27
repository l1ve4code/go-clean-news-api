package user

import (
	"go-clean-news-api/internal/domain/user"
	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *userStorage{
	return &userStorage{db: db}
}

func (bs *userStorage) GetOne(id int) *user.User {
	var user *user.User
	bs.db.Find(&user, id)
	return user
}
func (bs *userStorage) GetAll() []*user.User {
	return nil
}
func (bs *userStorage) Create(u *user.User) *user.User {
	bs.db.Create(u)
	return u
}
func (bs *userStorage) Delete(u *user.User) error {
	bs.db.Delete(u)
	return nil
}

func (bs *userStorage) FindByEmail(email string) *user.User{
	var user user.User
	bs.db.Where("email = ?", email).First(&user)
	return &user
}