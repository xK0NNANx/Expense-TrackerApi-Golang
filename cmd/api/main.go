package main

import (
	"expense-tracker/internal/handler"
	"expense-tracker/internal/repository"
	"expense-tracker/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	h := handler.New()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	db, err := repository.Connect(databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	log.Println("database connected successfully")

	repo := repository.NewExpRep(db)
	serv := service.NewExpenseService(repo)
	expenseHandler := handler.NewExpenseHandler(serv)

	r := chi.NewRouter()

	r.Get("/", h.Home)
	r.Get("/health", h.Health)

	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", expenseHandler.CreateExpense)
		r.Get("/", expenseHandler.GetExpenses)
	})

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
