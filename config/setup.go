package config

import (
	"database/sql"
	"fmt"
	"time"

	usersHandler "dating-apps/domain/users/handler"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"gopkg.in/gcfg.v1"
)

type Handler struct {
	Users *usersHandler.Handler
}

func LoadConfig() (*Config, error) {
	filename := "config.toml"

	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, filename)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitDatabase(cfg *Config) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	// PING
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set up Idle and Open Connections
	db.SetMaxIdleConns(cfg.Database.MaxIdleConnection)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnection)
	db.SetConnMaxIdleTime(time.Duration(cfg.Database.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.Database.MaxLifetimeConnection) * time.Second)

	return db, nil
}

func InitRedis(cfg *Config) *redis.Client {
	add := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     add,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	return client
}

func InitHandler(Users *usersHandler.Handler) *Handler {
	return &Handler{
		Users: Users,
	}
}
