package app

import (
	_ "github.com/AnkitDhawale/TodoListApp/docs"
	"github.com/AnkitDhawale/TodoListApp/handlers"
	"github.com/AnkitDhawale/TodoListApp/helpers"
	"github.com/AnkitDhawale/TodoListApp/middlewares"
	"github.com/AnkitDhawale/TodoListApp/repositories"
	"github.com/AnkitDhawale/TodoListApp/services"
	"github.com/go-chi/chi/v5"
	chiMW "github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

// @title ToDo App API
// @version 1.0
// @description This is a sample server for a todo list API.
// @host localhost:8888
// @BasePath /

var dbClient *pgx.Conn
var validate *validator.Validate

func Start() {
	dbClient = createDatabaseConnection()
	defer dbClient.Close()

	router := chi.NewRouter()
	router.Use(
		chiMW.Logger,
		chiMW.Recoverer,
	)

	// Swagger endpoint
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// Set a timeout middleware to limit request handling time
	//router.Use(chiMW.Timeout(500 * time.Millisecond)) // 500ms timeout for requests
	validate = validator.New(validator.WithRequiredStructEnabled())

	authRepo := repositories.NewAuthRepoDb(dbClient)
	authService := services.NewDefaultAuthService(authRepo)

	userResolver := helpers.NewUserResolver(dbClient)
	userRepo := repositories.NewUserRepoDb(dbClient)
	userService := services.NewDefaultUserService(userRepo)

	taskRepo := repositories.NewTaskRepoDb(dbClient)
	taskService := services.NewDefaultTaskService(taskRepo)

	th := handlers.TaskHandler{Service: taskService, Validatorr: validate}
	uh := handlers.UserHandler{Service: userService, AuthService: authService}

	router.Post("/todoapp/login", uh.Login)
	router.Post("/todoapp/signup", uh.SignUp)

	// private apis: needs auth token
	router.Route("/todoapp", func(r chi.Router) {
		r.Use(middlewares.TokenResolver(userResolver))

		// define user routes
		r.Get("/users/all", uh.GetAllUsers)
		r.Patch("/user-update", uh.UpdateUser)

		// define task routes
		r.Post("/tasks", th.CreateNewTask)
		r.Get("/tasks", th.GetAllTasksOfAUser)
		r.Put("/tasks/{id}", th.UpdateTask)
		r.Delete("/tasks/{id}", th.DeleteTask)
		r.Get("/tasks/view", th.ViewTasksWithFilter)
	})

	log.Println("Server started on localhost:8888...")
	log.Fatal(http.ListenAndServe(":8888", router))
}

// keeep it in repo
func createDatabaseConnection() *pgx.Conn {
	config := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOSTNAME"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("unable to establish db connection, err %v", err)
	}

	return conn
}