package composites

import "go-clean-news-api/internal/adapters/api/user"

type ControllersComposite struct {
	UserController user.Handler
}

func NewControllerComposite(userComposite *UserComposite) *ControllersComposite{
	return &ControllersComposite{UserController: userComposite.Handler}
}