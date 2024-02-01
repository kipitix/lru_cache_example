package psql

import (
	"context"
	"fmt"
	"lrucache/domain"
	"lrucache/domain/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type userRepositoryImpl struct {
	client *sqlx.DB
}

type PSQLCfg struct {
	// Data Source Name
	DSN string `arg:"--psql-dsn,env:PSQL_DSN" default:"host=localhost port=5432 user=postgres password=postgres dbname=lrucache_psql sslmode=disable"`
}

func NewUserRepository(cfg PSQLCfg) (repository.UserRepository, error) {
	db, err := sqlx.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}

	repo := userRepositoryImpl{client: db}
	if err = repo.client.Ping(); err != nil {
		return &repo, err
	}

	return &repo, nil
}

var _ repository.UserRepository = (*userRepositoryImpl)(nil)

func (r *userRepositoryImpl) InsertOrUpdateUser(ctx context.Context, user domain.User) error {
	query, args, err := prepareInsertOrUpdateUser(user)
	if err != nil {
		return fmt.Errorf("failed to prepare InsertOrUpdateUser query: %w", err)
	}
	_, err = r.client.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute InsertOrUpdateUser query: %w", err)
	}

	return nil
}

func (r *userRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	query, args, err := prepareFindUserByEmailQuery(email)
	if err != nil {
		return user, fmt.Errorf("failed to prepare FindUsers query: %w", err)
	}
	row := r.client.QueryRowxContext(ctx, query, args...)
	if err := row.StructScan(&user); err != nil {
		return user, fmt.Errorf("failed to parse row in FindUserByEmail: %w", err)
	}

	return user, err
}
