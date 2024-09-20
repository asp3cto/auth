package user

import (
	"context"
	"log"

	"github.com/asp3cto/auth/internal/converter"
	desc "github.com/asp3cto/auth/pkg/user_v1"
)

// Get ...
func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", user)

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
