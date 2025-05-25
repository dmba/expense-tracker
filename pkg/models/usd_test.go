package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToUSDConvertsFloatToUSD(t *testing.T) {
	result := ToUSD(15.75)

	assert.Equal(t, USD(1575), result)
}

func TestToUSDHandlesRounding(t *testing.T) {
	result := ToUSD(15.755)

	assert.Equal(t, USD(1576), result)
}

func TestFloat64ReturnsCorrectFloatValue(t *testing.T) {
	value := USD(1575)

	result := value.Float64()

	assert.Equal(t, 15.75, result)
}

func TestMultiplyReturnsCorrectResult(t *testing.T) {
	value := USD(1000)

	result := value.Multiply(1.5)

	assert.Equal(t, USD(1500), result)
}

func TestMultiplyHandlesZeroMultiplier(t *testing.T) {
	value := USD(1000)

	result := value.Multiply(0)

	assert.Equal(t, USD(0), result)
}

func TestStringFormatsUSDProperly(t *testing.T) {
	value := USD(1575)

	result := value.String()

	assert.Equal(t, "$15.75", result)
}

func TestMarshalCSVReturnsStringRepresentation(t *testing.T) {
	value := USD(1575)

	result, err := value.MarshalCSV()

	assert.NoError(t, err)
	assert.Equal(t, "1575", result)
}

func TestMarshalCSVHandlesNilValue(t *testing.T) {
	var value *USD

	result, err := value.MarshalCSV()

	assert.NoError(t, err)
	assert.Equal(t, "", result)
}

func TestUnmarshalCSVParsesValidString(t *testing.T) {
	var value USD

	err := value.UnmarshalCSV("1575")

	assert.NoError(t, err)
	assert.Equal(t, USD(1575), value)
}

func TestUnmarshalCSVHandlesEmptyString(t *testing.T) {
	var value USD

	err := value.UnmarshalCSV("")

	assert.NoError(t, err)
	assert.Equal(t, USD(0), value)
}

func TestUnmarshalCSVReturnsErrorForInvalidString(t *testing.T) {
	var value USD

	err := value.UnmarshalCSV("invalid")

	assert.Error(t, err)
}
