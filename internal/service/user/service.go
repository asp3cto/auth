package user

import (
	"github.com/asp3cto/auth/internal/client/db"
	"github.com/asp3cto/auth/internal/repository"
	"github.com/asp3cto/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService ...
func NewService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
