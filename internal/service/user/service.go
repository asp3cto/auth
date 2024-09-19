package user

import (
	"github.com/asp3cto/auth/internal/repository"
	"github.com/asp3cto/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

// NewService ...
func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}
