package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
	"expense-tracker/internal/repository"
)

type ExpenseService interface {
	AddExpense(dto.ExpenseDto) (*models.Expense, error)
	GetAllExpenses() ([]models.Expense, error)
}

type ExpenseServ struct {
	rep repository.ExpenseRepository
}

func NewExpenseService(rep repository.ExpenseRepository) *ExpenseServ {
	return &ExpenseServ{rep: rep}
}

func (serv *ExpenseServ) AddExpense(expDto dto.ExpenseDto) (*models.Expense, error) {
	expense := models.Expense{
		Title:  expDto.Title,
		Amount: expDto.Amount,
	}

	return serv.rep.Add(expense)
}

func (serv *ExpenseServ) GetAllExpenses() ([]models.Expense, error) {
	return serv.rep.GetAll()
}
