package composites

import (
	user3 "go-clean-news-api/internal/adapters/api/user"
	"go-clean-news-api/internal/adapters/db/user"
	user2 "go-clean-news-api/internal/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user2.Service
	Handler user3.Handler
}

func NewUserComposite(composite *MySQLComposite) *UserComposite{
	storage := user.NewStorage(composite.db)
	service := user2.NewService(storage)
	handler := user3.NewHandler(service)
	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}