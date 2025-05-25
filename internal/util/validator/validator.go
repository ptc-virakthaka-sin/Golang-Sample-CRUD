package validator

import (
	"learn-fiber/api/response"
	"reflect"
	"regexp"
	"strings"

	enlocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	englishNameRegex = regexp.MustCompile(`^[[:alpha:]]+(\s+[[:alpha:]]+)*$`)
	khmerNameRegex   = regexp.MustCompile(`^\p{Khmer}+(?:\s\p{Khmer}+)*$`)
)

type Validator struct {
	v   *validator.Validate
	uni *ut.UniversalTranslator
}

var V *Validator

func Init() {
	v := validator.New()
	locale := enlocale.New()
	uni := ut.New(locale, locale)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	_ = entranslations.RegisterDefaultTranslations(v, trans)
	_ = v.RegisterValidation("enname", englishNameValidation)
	_ = v.RegisterValidation("kmname", khmerNameValidation)
	v.RegisterTagNameFunc(omittedJSONFieldTagName)

	addTranslation(v, trans, "enname", "{0} must contain only english characters and allow whitespace between characters")
	addTranslation(v, trans, "kmname", "{0} must contain only khmer characters and allow whitespace between characters")

	V = &Validator{v: v, uni: uni}
}

func (vt *Validator) Valid(s interface{}) (hasError bool, errors []response.ValidationError) {
	errors = []response.ValidationError{}
	err := vt.v.Struct(s)
	if err != nil {
		hasError = true
		trans, _ := vt.uni.GetTranslator("en")
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.ValidationError{Message: err.Translate(trans), Field: err.Field()})
		}
	}
	return
}

func omittedJSONFieldTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func addTranslation(v *validator.Validate, trans ut.Translator, tag string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = v.RegisterTranslation(tag, trans, registerFn, transFn)
}

func englishNameValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	if v == "" {
		return true
	}
	return englishNameRegex.MatchString(v)
}

func khmerNameValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	if v == "" {
		return true
	}
	return khmerNameRegex.MatchString(v)
}
