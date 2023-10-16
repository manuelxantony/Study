package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const EMAIL_RX = "[a-z0-9]+@[a-z0-9]+\\.[a-z]{2,3}"

var EmailRX = regexp.MustCompile(EMAIL_RX)

type Validator struct{
	FieldErrors map[string]string
	NonFieldErrors []string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddFieldError(key, value string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = value
	}
}

func (v *Validator) AddNonFiledError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *Validator) CheckFiled(ok bool, key, value string) {
	if !ok {
		v.AddFieldError(key, value)
	}
}

func Matches(s string, rx *regexp.Regexp) bool {
	return rx.MatchString(s)
}

func IsNotBlank(s string)  bool {
	return  strings.TrimSpace(s) != ""
}

func IsMaxChars(s string, n int) bool {
	return utf8.RuneCountInString(s)  <= n
}

func IsMinChars(s string, n int) bool {
	return utf8.RuneCountInString(s) >= n
}

func IsPermittedValue[T comparable](value T, permittedValues ...T) bool {
	for _, pv := range permittedValues {
		if value == pv {
			return true
		}
	}
	return false
}