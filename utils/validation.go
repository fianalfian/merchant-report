package utils

import (
	"merchant-report/model"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	v        = validator.New()
	enTrans  = en.New()
	uni      = ut.New(enTrans, enTrans)
	trans, _ = uni.GetTranslator("en")
)

func NewValidator() (*validator.Validate, ut.Translator) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	en_translations.RegisterDefaultTranslations(v, trans)

	return v, trans
}

type Validator struct {
	Trans     ut.Translator
	Validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		msg := make(map[string]interface{})

		for _, e := range err.(validator.ValidationErrors) {
			msg[e.Field()] = map[string]string{"tag": e.Tag(), "message": e.Translate(v.Trans)}
		}

		return echo.NewHTTPError(http.StatusUnprocessableEntity, model.WebResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE_ENTITY",
			Data:   msg,
		})
	}
	return nil
}
