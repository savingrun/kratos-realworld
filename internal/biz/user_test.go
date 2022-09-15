package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := hashPassword("savingrun")
	spew.Dump(password)
}

func TestVerifyPassword(t *testing.T) {
	assert.New(t).True(verifyPassword("$2a$10$gekUq56rOED4Gm/PQlvvAug2wBpAnTMddG/M3vKBY5P3fpbxhgxwm",
		"savingrun"))
	assert.New(t).False(verifyPassword("$2a$10$gekUq56rOED4Gm/PQlvvAug2wBpAnTMddG/M3vKBY5P3fpbxhgxwm",
		"savingrun1"))
}
