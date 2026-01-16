package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewUser(t *testing.T) {
	user, err := NewUser("newuser@gmail.com", "jim@golang")
	assert.Nil(t, err)
	assert.NotNil(t, user.EncryptedPassword)
}

func TestUserPassword(t *testing.T) {
	pw := "jim@golang"
	user, err := NewUser("newuser@gmail.com", pw)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(pw))
	assert.False(t, user.ValidatePassword("jim@golang"))
}
