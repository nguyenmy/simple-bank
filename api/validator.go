package api

import (
	"fmt"
	"go-simple-bank/db/util"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validateCurrency validator.Func = func(fl validator.FieldLevel) bool {
	fmt.Println(fl.FieldName())
	if strings.EqualFold(fl.FieldName(), "Currency") {
		return util.IsCurrencySupported(fl.Field().String())
	}
	return false
}
