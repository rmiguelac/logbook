package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmiguelac/logbook/backend/internal/database"
	"github.com/rmiguelac/logbook/backend/internal/handlers"
	"github.com/rmiguelac/logbook/backend/internal/repositories"
	"github.com/rmiguelac/logbook/backend/pkg/auth"
	"github.com/rmiguelac/logbook/backend/pkg/config"
)

func main() {
	cfg := config.Load()

	// Database setup
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	authHandler := handlers.NewAuthHandler(userRepo, cfg.JWTSecret)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	// Router setup
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(auth.JWTMiddleware(cfg.JWTSecret))

	taskRouter := protected.PathPrefix("/tasks").Subrouter()
	taskRouter.HandleFunc("", taskHandler.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", taskHandler.GetTask).Methods("GET")

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
