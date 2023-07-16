package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
	"log"
	"os"
	"uji/database/aws"
	"uji/database/postgre"
	"uji/database/redis"
	"uji/database/redis/repository"
	"uji/helpers"
	userhandler "uji/users/delivery/http"
	userrepository "uji/users/repository/postgre"
	userusecase "uji/users/usecase"

	sosmedhandler "uji/social_media/delivery/http"
	sosmedrepository "uji/social_media/repository/postgre"
	sosmedusecase "uji/social_media/usecase"

	photohandler "uji/photo/delivery/http"
	photorepository "uji/photo/repository/postgre"
	photousecase "uji/photo/usecase"

	commenthandler "uji/comment/delivery/http"
	commentrepository "uji/comment/repository/postgre"
	commentusecase "uji/comment/usecase"
)

//func init() {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//}

func main() {
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	redisClient, err := redis.NewRedisClient(REDIS_HOST, REDIS_PASSWORD, 0)
	if err != nil {
		log.Fatalf("failed to create Redis client: %v", err)
	}

	svc := aws.InitS3()

	db, err := postgre.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	posgreDB, err := db.DB()
	err = posgreDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := posgreDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Recover())

	redisRepo := repository.NewRedisRepository(redisClient)

	//user ednpoint
	userRepo := userrepository.NewUserRepsitory(db, redisRepo)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userhandler.NewUserHandler(e, userUseCase, db)

	//sosmed endpoint
	sosmedRepo := sosmedrepository.NewSosmedRepsitory(db, redisRepo)
	sosmedUseCase := sosmedusecase.NewSosmedUseCase(sosmedRepo)
	sosmedhandler.NewSosmedHandler(e, sosmedUseCase, db)

	//photo endpoint
	photoRepo := photorepository.NewPhotoRepository(db, redisRepo)
	photoUseCase := photousecase.NewPhotoRepository(photoRepo)
	photohandler.NewPhotoHandler(e, photoUseCase, db, svc)

	//comment endpoint
	commentRepo := commentrepository.NewCommentRepository(db, redisRepo)
	commentUseCase := commentusecase.NewCommentUseCase(commentRepo)
	commenthandler.NewCommentUseCase(e, commentUseCase, db)

	e.Start(helpers.GetPort())

}
