package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo-go/auth"
	"todo-go/todo"

	authhttp "todo-go/auth/delivery/http"
	authmongo "todo-go/auth/repository/mongo"
	authUseCase "todo-go/auth/usecase"

	tdhttp "todo-go/todo/delivery/http"
	tdmongo "todo-go/todo/repository/mongo"
	tdusecase "todo-go/todo/usecase"
)

type App struct {
	httpServer *http.Server

	authUseCase auth.UseCase
	todoUseCase todo.UseCase
}

func NewApp() *App {
	db := initDb()

	userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	todoRepo := tdmongo.NewTodoRepository(db, viper.GetString("mongo.todo_collection"))

	return &App{
		authUseCase: authUseCase.NewAutoUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl")),

		todoUseCase: tdusecase.NewTodoUseCase(todoRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	authhttp.RegisterHTTPEndpoints(router, a.authUseCase)

	authMiddleware := authhttp.NewAutoMiddleware(a.authUseCase)
	api := router.Group("/api", authMiddleware)

	tdhttp.RegisterHTTPEndpoints(api, a.todoUseCase)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDb() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))

	if err != nil {
		log.Fatalf("error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}
