package user

import (
	"go-clean-news-api/internal/domain/entity"
)

type Storage interface {
	GetOne(id int) *entity.User
	GetAll() []*entity.User
	Create(user *entity.User) *entity.User
	Delete(user *entity.User) error
	FindByEmail(email string) *entity.User
}