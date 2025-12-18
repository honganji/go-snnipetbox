package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

// returns true if the FieldErrors map does not contain any entries
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// adds an error message to the FieldErrors map for a given key
func (v *Validator) AddFieldError(key, message string) {
	// lazy initialization of the map
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	// only add the error if an entry for the given key doesn't already exist
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// checks a boolean condition and adds an error message to the FieldErrors map if the condition is false
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// checks that a string is not blank (after trimming whitespace)
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// returns true if the provided string's length is less than or equal to the specified limit
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// returns true if the provided value is in the permittedValues slice
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
