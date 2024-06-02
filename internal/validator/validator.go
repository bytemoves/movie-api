package validator

import (
	"regexp"
	"slices"
)

// The movie title provided by the client is not empty and is not more than 500 bytes long.
// The movie year is not empty and is between 1888 and the current year.
// The movie runtime is not empty and is a positive integer.
// The movie has between one and five (unique) genres.



var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string] string
}

// helper creates new validator instance with an empty error map

func New() *Validator {
	return  &Validator{
		Errors: make(map[string]string),
	}
}

//validator returns true if the error map doeasnt contain ay entries

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

//Add error ads an error messafe to the map

func (v *Validator) AddError ( key , message string)  {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// check add an error message to the map only if a validation check os not ok

func(v *Validator) Check (ok bool, key, message string) {
	if !ok {
		v.AddError(key,message)
	}

}

// Generic function which returns true if a specific value is in a list of permitted
// values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
    return slices.Contains(permittedValues, value)
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
    return rx.MatchString(value)
}

// Generic function which returns true if all values in a slice are unique.
func Unique[T comparable](values []T) bool {
    uniqueValues := make(map[T]bool)

    for _, value := range values {
        uniqueValues[value] = true
    }

    return len(values) == len(uniqueValues)
}