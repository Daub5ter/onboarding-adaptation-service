package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/get_all", app.GetAll)
	mux.Put("/solve_instruction", app.SolveInstruction)
	mux.Post("/get_instruction", app.GetInstructionByID)
	mux.Post("/get_users_instructions", app.GetUsersInstructions)
	mux.Post("/add_instruction", app.AddInstruction)
	mux.Post("/add_users_instruction", app.AddUsersInstruction)
	mux.Post("/get_percent", app.GetPercent)

	return mux
}
