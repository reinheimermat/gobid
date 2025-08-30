package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/reinheimermat/gobid/internal/api"
	"github.com/reinheimermat/gobid/internal/services"
)

func main() {
	gob.Register(uuid.UUID{})

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBID_DATABASE_USER"),
		os.Getenv("GOBID_DATABASE_PASSWORD"),
		os.Getenv("GOBID_DATABASE_HOST"),
		os.Getenv("GOBID_DATABASE_PORT"),
		os.Getenv("GOBID_DATABASE_NAME"),
	))

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Failed to ping database:", err)
		return
	}

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: *services.NewUserService(pool),
		Sessions:    s,
	}

	api.BindRoutes()

	fmt.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", api.Router); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
