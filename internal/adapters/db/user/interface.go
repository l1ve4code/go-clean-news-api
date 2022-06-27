package user

import "go-clean-news-api/internal/domain/user"

type Storage interface {
	GetOne(id int) *user.User
	GetAll() []*user.User
	Create(user *user.User) *user.User
	Delete(user *user.User) error
	FindByEmail(email string) *user.User
}