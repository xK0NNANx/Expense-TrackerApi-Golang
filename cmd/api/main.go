package main

import (
	"expense-tracker/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	h := handler.New()
	r := chi.NewRouter()
	r.Get("/home", h.Home)
	r.Get("/health", h.Health)
	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", h.CreateExpense)
		r.Get("/{id}", h.GetExpenseByID)
		r.Get("/", h.GetExpenses)
		r.Delete("/{id}", h.DeleteExpense)
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
