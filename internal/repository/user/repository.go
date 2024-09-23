package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/asp3cto/auth/internal/client/db"
	serviceModel "github.com/asp3cto/auth/internal/model"
	"github.com/asp3cto/auth/internal/repository"
	"github.com/asp3cto/auth/internal/repository/user/converter"
	"github.com/asp3cto/auth/internal/repository/user/model"
	// "github.com/jackc/pgx/v4/pgxpool"
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
	db db.Client
}

// NewRepository ...
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

// Create ...
func (r *repo) Create(ctx context.Context, info *serviceModel.UserInfo) (int64, error) {
	builder := sq.Insert(tablename).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn).
		Values(info.Name, info.Email).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
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

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user model.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
