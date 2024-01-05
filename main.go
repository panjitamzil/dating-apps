package main

import (
	"dating-apps/config"
	usersHandler "dating-apps/domain/users/handler"
	usersRepository "dating-apps/domain/users/repository"
	usersUsecase "dating-apps/domain/users/usecase"
	"dating-apps/route"
	"fmt"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Connect to database
	db, err := config.InitDatabase(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Initialized")

	// Init Redis
	redis := config.InitRedis(cfg)
	_, err = redis.Ping(redis.Context()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis Initialized")

	usersRepository := usersRepository.NewRepo(db, redis)
	usersUsecase := usersUsecase.NewService(usersRepository)
	usersHandler := usersHandler.NewHandler(usersUsecase)

	// Init Handler
	handler := config.InitHandler(usersHandler)

	r := route.Routes(handler)
	fmt.Printf("%s is running ... \n", cfg.Application.ServiceName)
	r.Run(cfg.Application.Port)
}
