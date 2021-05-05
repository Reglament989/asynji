package validators

import (
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/Reglament989/asynji/pkgs/asynji/validators/errors"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type MultiValidation struct {
	validators []Validator
}

func (e *MultiValidation) Add(v Validator) {
	e.validators = append(e.validators, v)
}

func (e *MultiValidation) Result() []error {
	allErrors := []error{}
	for idx := range e.validators {
		if len(e.validators[idx].Errors) > 0 {
			allErrors = append(allErrors, e.validators[idx].Errors...)
		}
	}
	if len(allErrors) > 0 {
		return allErrors
	}
	return nil
}

type Validator struct {
	Data    string
	Success bool
	Errors  []error
	Name    string
}

func (d *Validator) Length(minLength int, maxLength int) {
	if len(d.Data) > minLength && len(d.Data) < maxLength {
		d.Success = true
		return
	}
	d.Errors = append(d.Errors, &errors.LengthError{MinLength: minLength, MaxLength: maxLength, ObjName: d.Name})
}

func (d *Validator) Result() []error {
	if len(d.Errors) > 0 {
		return d.Errors
	}
	return nil
}

func (d *Validator) IsEmail() {
	if len(d.Data) < 3 && len(d.Data) > 254 {
		d.Success = false
		d.Errors = append(d.Errors, &errors.IsEmailError{})
	}
	if !emailRegex.MatchString(d.Data) {
		d.Success = false
		d.Errors = append(d.Errors, &errors.IsEmailError{})
	}
	parts := strings.Split(d.Data, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		d.Success = false
		d.Errors = append(d.Errors, &errors.IsEmailError{})
	}
	d.Success = true
}

func (d *Validator) IsUrl() {
	_, err := url.ParseRequestURI(d.Data)
	if err != nil {
		d.Success = false
		d.Errors = append(d.Errors, &errors.IsEmailError{})
	}
}

// type Validator interface {
// 	Result() bool
// }
