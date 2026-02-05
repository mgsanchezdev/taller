package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	apimw "workshop-platform/backend/internal/api/middleware"
	"workshop-platform/backend/internal/api/handlers"
	"workshop-platform/backend/internal/api/routes"
	"workshop-platform/backend/internal/auth"
	"workshop-platform/backend/internal/infrastructure/postgres"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-in-production"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/workshop_platform?sslmode=disable"
	}

	db, err := postgres.New(dbURL)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	jwt := auth.NewJWT(jwtSecret)
	userRepo := postgres.NewUserRepo(db)
	vehicleRepo := postgres.NewVehicleRepo(db)

	d := &routes.Deps{
		AuthHandler:     handlers.NewAuthHandler(jwt, userRepo),
		VehiclesHandler: handlers.NewVehiclesHandler(vehicleRepo),
		JWT:             jwt,
	}

	// Ensure auth middleware is used (already in routes)
	_ = apimw.Auth(jwt)

	router := routes.Router(d)
	log.Printf("API listening on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
