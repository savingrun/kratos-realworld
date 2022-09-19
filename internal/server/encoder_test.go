package server

import (
	"encoding/json"
	"kratos-realworld/internal/errors"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHttpStruct(t *testing.T) {
	a := &errors.HttpError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be empty"}
	marshal, err := json.Marshal(a)
	assert.NoError(t, err)
	spew.Dump(string(marshal))
}
