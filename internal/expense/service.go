package expense

import (
	"errors"
	"github.com/dmba/expense-tracker/pkg/fs"
	"github.com/dmba/expense-tracker/pkg/models"
	"github.com/dmba/expense-tracker/pkg/utils"
	"slices"
	"time"
)

var (
	ErrExpenseNotFound = errors.New("expense not found")
)

type Storage[T any] interface {
	Read() (T, error)
	Write(data T) error
}

type Service struct {
	file Storage[[]models.Expense]
}

func NewService(filepath string) *Service {
	file := fs.NewCsv[[]models.Expense](filepath)
	return &Service{
		file: file,
	}
}

func (s *Service) AddExpense(description string, amount float64) (*models.Expense, error) {
	expenses, err := s.file.Read()
	if err != nil {
		return nil, err
	}

	id := utils.NextId(expenses, func(e models.Expense) int {
		return e.ID
	})
	newExpense := models.Expense{
		ID:          id,
		Description: description,
		Amount:      models.ToUSD(amount),
		Date:        time.Now(),
	}
	expenses = append(expenses, newExpense)

	if err := s.file.Write(expenses); err != nil {
		return nil, err
	}

	return &newExpense, nil
}

func (s *Service) DeleteExpense(id int) error {
	expenses, err := s.file.Read()
	if err != nil {
		return err
	}

	modified := slices.DeleteFunc(expenses, func(e models.Expense) bool {
		return e.ID == id
	})

	if len(expenses) == len(modified) {
		return ErrExpenseNotFound
	}

	if err := s.file.Write(modified); err != nil {
		return err
	}

	return nil
}

func (s *Service) ListExpenses() ([]models.Expense, error) {
	expenses, err := s.file.Read()
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *Service) Summary() (models.USD, error) {
	expenses, err := s.file.Read()
	if err != nil {
		return -1, err
	}
	var total models.USD
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total, nil
}

func (s *Service) SummaryByMonth(month time.Month) (models.USD, error) {
	expenses, err := s.file.Read()
	if err != nil {
		return -1, err
	}
	var total models.USD
	currentYear := time.Now().Year()
	for _, expense := range expenses {
		if currentYear == expense.Date.Year() && expense.Date.Month() == month {
			total += expense.Amount
		}
	}
	return total, nil
}
