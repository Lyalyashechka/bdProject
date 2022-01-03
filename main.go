package main

import (
	"database/sql"
	"fmt"
	forumHandler "github.com/Lyalyashechka/bdProject/app/forum/handler"
	forumRepository "github.com/Lyalyashechka/bdProject/app/forum/repository"
	forumUC "github.com/Lyalyashechka/bdProject/app/forum/usecase"
	threadHandler "github.com/Lyalyashechka/bdProject/app/thread/handler"
	threadRepository "github.com/Lyalyashechka/bdProject/app/thread/repository"
	threadUC "github.com/Lyalyashechka/bdProject/app/thread/usecase"
	"github.com/Lyalyashechka/bdProject/app/tools"
	userHandler "github.com/Lyalyashechka/bdProject/app/user/handler"
	userRepository "github.com/Lyalyashechka/bdProject/app/user/repository"
	userUC "github.com/Lyalyashechka/bdProject/app/user/usecase"
	validator "github.com/go-playground/validator"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/labstack/echo/v4"
	"log"
)

var (
	router = echo.New()
)

func main() {
	db, err := GetPostgres()
	if err != nil {
		log.Fatal(err)
	}

	userHandler := userHandler.NewHandler(userUC.NewUseCase(
		userRepository.NewRepository(db)))
	forumHandler := forumHandler.NewHandler(forumUC.NewUseCase(
		forumRepository.NewRepository(db), threadRepository.NewRepository(db)))
	threadHandler := threadHandler.NewHandler(threadUC.NewUseCase(
		threadRepository.NewRepository(db)))

	validator := validator.New()
	router.Validator = tools.NewCustomValidator(validator)

	router.POST("user/:nickname/create", userHandler.SignUpUser)
	router.GET("user/:nickname/profile", userHandler.GetUser)
	router.POST("user/:nickname/profile", userHandler.UpdateUser)
	router.POST("forum/create", forumHandler.CreateForum)
	router.GET("forum/:slug/details", forumHandler.GetForumDetails)
	router.POST("forum/:slug/create", forumHandler.CreateThread)
	router.GET("forum/:slug/users", forumHandler.GetUsersForum)
	router.GET("forum/:slug/threads", forumHandler.GetForumThreads)
	router.POST("thread/:slug_or_id/create", threadHandler.CreatePosts)
	if err := router.Start("127.0.0.1:5000"); err != nil {
		log.Fatal(err)
	}
}

func GetPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		"eugeniy", "postgres",
		"docker", "127.0.0.1",
		"5432")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	return db, nil
}
