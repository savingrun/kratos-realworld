package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("savingrun", "saving")
	spew.Dump(token)
}
