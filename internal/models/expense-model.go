package models

type Expense struct {
	ID     int     `db:"id" json:"id"`
	Title  string  `db:"title" json:"title"`
	Amount float64 `db:"amount" json:"amount"`
}
