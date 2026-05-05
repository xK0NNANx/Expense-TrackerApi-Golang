package repository

import (
	"expense-tracker/internal/models"

	"github.com/jmoiron/sqlx"
)

type ExpenseRepository interface {
	Add(models.Expense) (*models.Expense, error)
	GetAll() ([]models.Expense, error)
}

type ExpenseRep struct {
	db *sqlx.DB
}

func NewExpRep(db *sqlx.DB) *ExpenseRep {
	return &ExpenseRep{db: db}
}

func (exRep *ExpenseRep) Add(exp models.Expense) (*models.Expense, error) {
	query := `
		INSERT INTO expenses (title, amount)
		VALUES ($1, $2)
		RETURNING id, title, amount
	`

	err := exRep.db.QueryRowx(query, exp.Title, exp.Amount).StructScan(&exp)
	if err != nil {
		return nil, err
	}

	return &exp, nil
}

func (exRep *ExpenseRep) GetAll() ([]models.Expense, error) {
	var expenses []models.Expense

	query := `
		SELECT id, title, amount
		FROM expenses
		ORDER BY id
	`

	err := exRep.db.Select(&expenses, query)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
