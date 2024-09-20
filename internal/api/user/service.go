package user

import (
	"github.com/asp3cto/auth/internal/service"
	desc "github.com/asp3cto/auth/pkg/user_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation ...
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
