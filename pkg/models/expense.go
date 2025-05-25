package models

import (
	"time"
)

type Expense struct {
	ID          int       `csv:"id"`
	Description string    `csv:"description"`
	Amount      USD       `csv:"amount"`
	Date        time.Time `csv:"date"`
}
