package db_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ibiscum/Test-Driven-Development-in-Go/chapter06/db"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	ps := db.NewPostingService()
	b := db.Book{
		ID: uuid.New().String(),
	}
	err := ps.NewOrder(b)
	assert.Nil(t, err)
}
