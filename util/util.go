package util

import "strings"

type JError struct {
	Error string `json:"error"`
}

func NewJError(err error) JError {
	jerr := JError{"generic error"}
	if err != nil {
		jerr.Error = err.Error()
	}
	return jerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
