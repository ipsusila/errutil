package errutil

import (
	"bytes"
	"fmt"
)

type Errors struct {
	prefix string
	errs   []error
}

// New return errors
func New() *Errors {
	return NewWithPrefix("")
}

// NewWithPrefix return errors with prefix
func NewWithPrefix(prefix string) *Errors {
	return &Errors{
		prefix: prefix,
	}
}

// Error interface
func (e *Errors) Error() string {
	buf := bytes.Buffer{}
	for _, err := range e.errs {
		buf.WriteString(fmt.Sprintln(e.prefix + err.Error()))
	}
	return buf.String()
}

// AsError return nil if no error, otherwise return self
func (e *Errors) AsError() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

// Append new error to the list
func (e *Errors) Append(err error) *Errors {
	if err != nil {
		e.errs = append(e.errs, err)
	}
	return e
}

// Items return list of errors
func (e *Errors) Items() []error {
	return e.errs
}
