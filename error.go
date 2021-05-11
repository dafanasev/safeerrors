package safeerrors

import (
	"strings"
	"sync"
)

type Error struct {
	mu   sync.Mutex
	errs []error
}

func Empty() *Error {
	return &Error{
		mu:   sync.Mutex{},
		errs: []error{},
	}
}

func (e *Error) Set(err error) {
	e.mu.Lock()
	e.errs = []error{err}
	e.mu.Unlock()
}

func (e *Error) Append(err error) {
	e.mu.Lock()
	e.errs = append(e.errs, err)
	e.mu.Unlock()
}

func (e *Error) Error() string {
	e.mu.Lock()
	defer e.mu.Unlock()

	var errStrs []string
	for _, err := range e.errs {
		errStrs = append(errStrs, err.Error())
	}

	return strings.Join(errStrs, "\n")
}