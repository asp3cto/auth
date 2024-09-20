package user

import (
	"context"
	"log"

	"github.com/asp3cto/auth/internal/converter"
	desc "github.com/asp3cto/auth/pkg/user_v1"
)

// Create ...
func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("Inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
