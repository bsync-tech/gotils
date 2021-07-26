package error

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultErrorIfNil(t *testing.T) {
	assert.Equal(t, DefaultErrorIfNil(nil, "Cool"), "Cool")
	assert.Equal(t, DefaultErrorIfNil(errors.New("Oops"), "Cool"), "Oops")
}
