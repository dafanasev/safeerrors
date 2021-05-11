package safeerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := Empty()
	err.Set(errors.New("first error"))
	assert.EqualError(t, err, "first error")
	err.Append(errors.New("second error"))
	assert.EqualError(t, err, "first error\nsecond error")
	err.Append(errors.New("third error"))
	assert.EqualError(t, err, "first error\nsecond error\nthird error")
	err.Set(errors.New("fourth error"))
	assert.EqualError(t, err, "fourth error")
}
