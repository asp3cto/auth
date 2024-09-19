package repository

import (
	"context"

	"github.com/asp3cto/auth/internal/model"
)

// UserRepository ...
type UserRepository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
