package user

import (
	"context"

	"github.com/asp3cto/auth/internal/model"
)

// Get ...
func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
