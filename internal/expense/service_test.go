package expense

import (
	"errors"
	"github.com/dmba/expense-tracker/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockFile[T any] struct {
	mock.Mock
}

func (m *MockFile[T]) Read() (T, error) {
	args := m.Called()
	var zero T
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}
	return args.Get(0).(T), args.Error(1)
}

func (m *MockFile[T]) Write(data T) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestAddExpenseAddsNewExpense(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}

	mockFile.On("Read").Return([]models.Expense{}, nil)
	mockFile.On("Write", mock.Anything).Return(nil)

	expense, err := service.AddExpense("Lunch", 15.50)

	assert.NoError(t, err)
	assert.NotNil(t, expense)
	assert.Equal(t, "Lunch", expense.Description)
	assert.Equal(t, models.ToUSD(15.50), expense.Amount)
	mockFile.AssertCalled(t, "Write", mock.Anything)
}

func TestAddExpenseHandlesFileReadError(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}

	mockFile.On("Read").Return(nil, errors.New("read error"))

	expense, err := service.AddExpense("Lunch", 15.50)

	assert.Error(t, err)
	assert.Nil(t, expense)
}

func TestDeleteExpenseDeletesExistingExpense(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}

	expenses := []models.Expense{
		{ID: 1, Description: "Lunch", Amount: models.ToUSD(15.50)},
	}
	mockFile.On("Read").Return(expenses, nil)
	mockFile.On("Write", mock.Anything).Return(nil)

	err := service.DeleteExpense(1)

	assert.NoError(t, err)
	mockFile.AssertCalled(t, "Write", []models.Expense{})
}

func TestDeleteExpenseReturnsErrorIfNotFound(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}

	expenses := []models.Expense{
		{ID: 1, Description: "Lunch", Amount: models.ToUSD(15.50)},
	}
	mockFile.On("Read").Return(expenses, nil)

	err := service.DeleteExpense(2)

	assert.ErrorIs(t, err, ErrExpenseNotFound)
}

func TestSummaryCalculatesTotalAmount(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}

	expenses := []models.Expense{
		{ID: 1, Amount: models.ToUSD(10.00)},
		{ID: 2, Amount: models.ToUSD(20.00)},
	}
	mockFile.On("Read").Return(expenses, nil)

	total, err := service.Summary()

	assert.NoError(t, err)
	assert.Equal(t, models.ToUSD(30.00), total)
}

func TestSummaryByMonthFiltersExpensesByMonth(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}
	now := time.Now()

	expenses := []models.Expense{
		{ID: 1, Amount: models.ToUSD(10.00), Date: time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Amount: models.ToUSD(20.00), Date: time.Date(now.Year(), time.February, 1, 0, 0, 0, 0, time.UTC)},
	}
	mockFile.On("Read").Return(expenses, nil)

	total, err := service.SummaryByMonth(time.January)

	assert.NoError(t, err)
	assert.Equal(t, models.ToUSD(10.00), total)
}

func TestSummaryByMonthIgnoresExpensesFromOtherYears(t *testing.T) {
	mockFile := &MockFile[[]models.Expense]{}
	service := &Service{file: mockFile}
	now := time.Now()

	expenses := []models.Expense{
		{ID: 1, Amount: models.ToUSD(10.00), Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Amount: models.ToUSD(20.00), Date: time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
	mockFile.On("Read").Return(expenses, nil)

	total, err := service.SummaryByMonth(time.January)

	assert.NoError(t, err)
	assert.Equal(t, models.ToUSD(20.00), total)
}
