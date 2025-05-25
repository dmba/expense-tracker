package fs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testType struct {
	ID          int     `csv:"ID"`
	Description string  `csv:"Description"`
	Amount      float64 `csv:"Amount"`
}

func createTempFile(t *testing.T, content string) (string, func()) {
	mockFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := mockFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	mockFile.Close()
	return mockFile.Name(), func() {
		os.Remove(mockFile.Name())
	}
}

func TestReadReturnsDataFromFile(t *testing.T) {
	fileName, cleanup := createTempFile(t, "ID,Description,Amount\n1,Lunch,15.50\n")
	defer cleanup()

	file := NewCsv[[]testType](fileName)

	data, err := file.Read()

	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, 1, data[0].ID)
	assert.Equal(t, "Lunch", data[0].Description)
	assert.Equal(t, 15.50, data[0].Amount)
}

func TestReadHandlesEmptyFile(t *testing.T) {
	fileName, cleanup := createTempFile(t, "")
	defer cleanup()

	file := NewCsv[[]testType](fileName)

	data, err := file.Read()

	assert.NoError(t, err)
	assert.Empty(t, data)
}

func TestWriteSavesDataToFile(t *testing.T) {
	fileName, cleanup := createTempFile(t, "")
	defer cleanup()

	file := NewCsv[[]testType](fileName)

	data := []testType{
		{ID: 1, Description: "Lunch", Amount: 15.50},
	}

	err := file.Write(data)

	assert.NoError(t, err)

	readData, _ := file.Read()
	assert.Len(t, readData, 1)
	assert.Equal(t, data[0], readData[0])
}
