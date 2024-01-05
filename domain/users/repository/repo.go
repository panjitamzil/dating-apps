package repository

import (
	"database/sql"
	"dating-apps/domain/users"

	"github.com/go-redis/redis/v8"
)

type Repo struct {
	db    *sql.DB
	redis *redis.Client
}

func NewRepo(db *sql.DB, redis *redis.Client) users.RepoInterface {
	return &Repo{
		db:    db,
		redis: redis,
	}
}
