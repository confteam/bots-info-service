package utils

import (
	"github.com/teris-io/shortid"
)

func GenerateCode() (string, error) {
	code, err := shortid.Generate()
	return code, err
}
