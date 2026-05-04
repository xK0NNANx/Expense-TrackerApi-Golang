package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Expense struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

var expenses []Expense

func (h *Handler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if expense.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	if expense.Amount <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		return
	}
	expense.ID = len(expenses) + 1
	expenses = append(expenses, expense)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(expense); err != nil {
		http.Error(w, "faile to encode", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		http.Error(w, "faile to encode", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetExpenseByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, expense := range expenses {
		if expense.ID == id {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(expense); err != nil {
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "expense not found", http.StatusNotFound)
}

func (h *Handler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"message":"expense deleted"}`))
			return
		}
	}
	http.Error(w, "expense not found", http.StatusNotFound)
}
