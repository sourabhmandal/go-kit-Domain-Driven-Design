package strconv

import (
	"errors"
	"strings"
)

// interface that defines methods to be
// exposed by StringService
type IStringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// imported in api package
// StringService implements IStringService
type StringService struct{}

// core logic for converting normal string 
// to uppercase sting
func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errEmpty
	}
	return strings.ToUpper(s), nil
}

// core logic to Count
// the number of charachters in a string
func (StringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var errEmpty = errors.New("empty string")