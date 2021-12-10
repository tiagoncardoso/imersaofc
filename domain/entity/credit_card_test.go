package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("999999999999999999", "James Cameron", 12, 2023, 333)
	assert.Equal(t, "Invalid credit card number", err.Error())

	_, err = NewCreditCard("5140231276348928", "James Cameron", 12, 2023, 333)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5140231276348928", "James Cameron", 14, 2023, 333)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("5140231276348928", "James Cameron", 2, 2023, 333)
	assert.Nil(t, err)
}

func (cc *CreditCard) TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0).Year()
	nextYear := time.Now().AddDate(1, 0, 0).Year()

	_, err := NewCreditCard("5140231276348928", "James Cameron", 2, lastYear, 333)
	assert.Equal(t, "Invalid expiration year", err.Error())

	_, err = NewCreditCard("5140231276348928", "James Cameron", 2, nextYear, 333)
	assert.Nil(t, err)
}
