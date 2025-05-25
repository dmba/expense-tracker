package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextSeqIdReturnsOneForEmptySlice(t *testing.T) {
	var data []struct{}
	idFunc := func(_ struct{}) int { return 0 }

	result := NextId(data, idFunc)

	assert.Equal(t, 1, result)
}

func TestNextSeqIdReturnsNextIdForNonEmptySlice(t *testing.T) {
	data := []struct {
		ID int
	}{
		{ID: 1},
		{ID: 3},
		{ID: 2},
	}
	idFunc := func(item struct{ ID int }) int { return item.ID }

	result := NextId(data, idFunc)

	assert.Equal(t, 4, result)
}

func TestNextSeqIdHandlesNegativeIds(t *testing.T) {
	data := []struct {
		ID int
	}{
		{ID: -5},
		{ID: -1},
		{ID: -3},
	}
	idFunc := func(item struct{ ID int }) int { return item.ID }

	result := NextId(data, idFunc)

	assert.Equal(t, 0, result)
}

func TestNextSeqIdHandlesSingleElementSlice(t *testing.T) {
	data := []struct {
		ID int
	}{
		{ID: 42},
	}
	idFunc := func(item struct{ ID int }) int { return item.ID }

	result := NextId(data, idFunc)

	assert.Equal(t, 43, result)
}
