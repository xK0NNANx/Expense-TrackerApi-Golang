package handler

import (
	"encoding/json"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/service"
	"net/http"
)

type ExpenseHandler struct {
	svc service.ExpenseService
}

func NewExpenseHandler(svc service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{svc: svc}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var expDto dto.ExpenseDto
	if err := json.NewDecoder(r.Body).Decode(&expDto); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if expDto.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	if expDto.Amount <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		return
	}

	expense, err := h.svc.AddExpense(expDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(expense); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ExpenseHandler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.svc.GetAllExpenses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
