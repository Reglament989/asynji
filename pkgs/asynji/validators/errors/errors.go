package errors

import "fmt"

type LengthError struct {
	MinLength int
	MaxLength int
	ObjName   string
}

func (m *LengthError) Error() string {
	return fmt.Sprintf("Length of %v must be a longer than %d and shorten than %d", m.ObjName, m.MinLength, m.MaxLength)
}

type IsEmailError struct {
}

func (m *IsEmailError) Error() string {
	return "This email is incorrect"
}

type IsUrlError struct {
}

func (m *IsUrlError) Error() string {
	return "This url is incorrect"
}
