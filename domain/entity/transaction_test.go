package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionIsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900

	assert.Nil(t, transaction.IsValid())
}

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 9000

	err := transaction.IsValid()

	assert.Equal(t, "Invalid transaction amount", err.Error())
}
