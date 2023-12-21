package helper

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var lock = &sync.Mutex{}
var validate *validator.Validate

func GetValidator() *validator.Validate {
	lock.Lock()
	defer lock.Unlock()

	if validate == nil {
		validate = validator.New()

		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	return validate
}
