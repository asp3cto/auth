package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	serviceModel "github.com/asp3cto/auth/internal/model"
	"github.com/asp3cto/auth/internal/repository"
	"github.com/asp3cto/auth/internal/repository/user/converter"
	"github.com/asp3cto/auth/internal/repository/user/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tablename       = `"user"`
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

// NewRepository ...
func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

// Create ...
func (r *repo) Create(ctx context.Context, info *serviceModel.UserInfo) (int64, error) {
	builder := sq.Insert(tablename).PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn).
		Values(info.Name, info.Email).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get ...
func (r *repo) Get(ctx context.Context, id int64) (*serviceModel.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tablename).Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
