package test

import (
	"strconv"
	"testing"

	c "github.com/garciaolais/challenge-backend/challengelib"
	util "github.com/garciaolais/challenge-backend/util"

	"github.com/stretchr/testify/assert"
)

func TestDivideByThree(t *testing.T) {
	data := []c.Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

	dividends := []int{3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39, 42, 48}
	for _, dividend := range dividends {
		message, ok := c.GetMessage(dividend, data)
		assert.True(t, ok, "Expected true")
		assert.Equal(t, "Linio", message, "Expected [Linio]")
	}
}

func TestDivideByFive(t *testing.T) {
	data := []c.Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

	dividends := []int{5, 10, 20, 25, 35, 40, 50, 55, 65, 70, 80, 85, 95, 100}
	for _, dividend := range dividends {
		message, ok := c.GetMessage(dividend, data)
		assert.True(t, ok, "Expected true")
		assert.Equal(t, "IT", message, "Expected [IT]")
	}
}

func TestDivideByFiveOrThree(t *testing.T) {
	data := []c.Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

	dividends := []int{15, 30, 45, 60, 75, 90}
	for _, dividend := range dividends {
		message, ok := c.GetMessage(dividend, data)
		assert.True(t, ok, "Expected true")
		assert.Equal(t, "Linianos", message, "Expected [Linianos]")
	}
}

func TestRun(t *testing.T) {
	data := []c.Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

	results, err := util.ReadLines("results.txt")
	assert.NoError(t, err, "Expected no error")

	for i, result := range results {
		_, err := strconv.Atoi(result)
		if err != nil {
			message, ok := c.GetMessage(i+1, data)
			assert.True(t, ok, "Expected true")
			assert.Equalf(t, result, message, "Expected [%s]", result)
		} else {
			resultActual := strconv.Itoa(i + 1)
			assert.Equalf(t, result, resultActual, "Expected [%s]", result)
		}
	}
}

func TestGetOnlyInts(t *testing.T) {
	data := []c.Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

	results, err := util.ReadLines("results-int.txt")
	assert.NoError(t, err, "Expected no error")

	for _, result := range results {
		ressultExpected, err := strconv.Atoi(result)
		assert.NoError(t, err, "Expected no error")
		message, ok := c.GetMessage(ressultExpected, data)
		assert.False(t, ok, "Expected false")
		assert.Equal(t, "", message, "Expected empty string")
	}
}
