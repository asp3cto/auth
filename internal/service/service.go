package service

import (
	"context"

	"github.com/asp3cto/auth/internal/model"
)

// UserService ...
type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
