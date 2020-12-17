package ru

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterDefaultTranslations registers a set of default translations
// for all built in tag's in validator; you may add your own as desired.
func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {

	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "required",
			translation: "{0} обязательное поле",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "Поле {0} должно быть длиной в {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "Поле {0} должно быть равно {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "Поле {0} должно содержать {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-items", fld, c)

				default:
					t, err = ut.T("len-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("min-string", "Поле {0} должно содержать минимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "Поле {0} должно быть больше или равно {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "Поле {0} должно содержать минимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}
				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", fld, c)

				default:
					t, err = ut.T("min-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("max-string", "Поле {0} должно содержать максимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "Поле {0} должно быть меньше или равно {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "Поле {0} должно содержать максимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", fld, c)

				default:
					t, err = ut.T("max-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0} не равен {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "Поле {0} должно быть не равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lt-string", "Поле {0} должно иметь менее {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "Поле {0} должно быть менее {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "Поле {0} должно содержать менее {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} должно быть меньше текущей даты и времени", false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-items", fld, c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lt-datetime", fld)

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lte-string", "Поле {0} должно содержать максимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "Поле {0} должно быть менее или равно {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "Поле {0} должно содержать максимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} должно быть меньше или равно текущей дате и времени", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-items", fld, c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lte-datetime", fld)

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gt-string", "Поле {0} должно быть длиннее {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "Поле {0} должно быть больше {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "Поле {0} должно содержать более {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} должна быть позже текущего момента", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-items", fld, c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gt-datetime", fld)

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gte-string", "Поле {0} должно содержать минимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} символ", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} символы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} символа", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} символов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "Поле {0} должно быть больше или равно {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "Поле {0} должно содержать минимум {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} элемент", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} элементы", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} элемента", locales.PluralRuleFew, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} элементов", locales.PluralRuleMany, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} должна быть позже или равна текущему моменту", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", fld, c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-items", fld, c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gte-datetime", fld)

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", fld, ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "Поле {0} должно быть равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "Поле {0} должно быть равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0} не должен быть равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "Поле {0} должно быть больше {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "Поле {0} должно быть больше или равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "Поле {0} должно быть менее {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "Поле {0} должно быть менее или равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "Поле {0} не должен быть равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "Поле {0} должно быть больше {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "Поле {0} должно быть больше или равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "Поле {0} должно быть менее {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "Поле {0} должно быть менее или равно {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "Поле {0} должно содержать только буквы",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "alphanum",
			translation: "Поле {0} должно содержать только буквы и цифры",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "numeric",
			translation: "Поле {0} должно быть цифровым значением",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "number",
			translation: "Поле {0} должно быть цифрой",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "hexadecimal",
			translation: "Поле {0} должно быть шестнадцатеричной строкой",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "hexcolor",
			translation: "Поле {0} должно быть HEX цветом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "rgb",
			translation: "Поле {0} должно быть RGB цветом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "rgba",
			translation: "Поле {0} должно быть RGBA цветом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "hsl",
			translation: "Поле {0} должно быть HSL цветом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "hsla",
			translation: "Поле {0} должно быть HSLA цветом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "e164",
			translation: "Поле {0} должно быть E.164 formatted phone number",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "email",
			translation: "Поле {0} должно быть email адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "url",
			translation: "Поле {0} должно быть URL",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "uri",
			translation: "Поле {0} должно быть URI",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "base64",
			translation: "Поле {0} должно быть Base64 строкой",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "contains",
			translation: "Поле {0} должно содержать текст '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "Поле {0} должно содержать минимум один из символов '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "Поле {0} не должно содержать текст '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "Поле {0} не должно содержать символы '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "Поле {0} не должно содержать '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				t, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "Поле {0} должно быть ISBN номером",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "isbn10",
			translation: "Поле {0} должно быть ISBN-10 номером",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "isbn13",
			translation: "Поле {0} должно быть ISBN-13 номером",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "uuid",
			translation: "Поле {0} должно быть UUID",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "uuid3",
			translation: "Поле {0} должно быть UUID 3 версии",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "uuid4",
			translation: "Поле {0} должно быть UUID 4 версии",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "uuid5",
			translation: "Поле {0} должно быть UUID 5 версии",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ascii",
			translation: "Поле {0} должно содержать только ascii символы",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "printascii",
			translation: "Поле {0} должно содержать только доступные для печати ascii символы",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "multibyte",
			translation: "Поле {0} должно содержать мультибайтные символы",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "datauri",
			translation: "Поле {0} должно содержать Data URI",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "latitude",
			translation: "Поле {0} должно содержать координаты широты",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "longitude",
			translation: "Поле {0} должно содержать координаты долготы",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ssn",
			translation: "Поле {0} должно быть SSN номером",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ipv4",
			translation: "Поле {0} должно быть IPv4 адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ipv6",
			translation: "Поле {0} должно быть IPv6 адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ip",
			translation: "Поле {0} должно быть IP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "cidr",
			translation: "Поле {0} должно содержать CIDR обозначения",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "cidrv4",
			translation: "Поле {0} должно содержать CIDR обозначения для IPv4 адреса",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "cidrv6",
			translation: "Поле {0} должно содержать CIDR обозначения для IPv6 адреса",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "tcp_addr",
			translation: "Поле {0} должно быть TCP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "tcp4_addr",
			translation: "Поле {0} должно быть IPv4 TCP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "tcp6_addr",
			translation: "Поле {0} должно быть IPv6 TCP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "udp_addr",
			translation: "Поле {0} должно быть UDP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "udp4_addr",
			translation: "Поле {0} должно быть IPv4 UDP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "udp6_addr",
			translation: "Поле {0} должно быть IPv6 UDP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ip_addr",
			translation: "Поле {0} должно быть распознаваемым IP адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ip4_addr",
			translation: "Поле {0} должно быть распознаваемым IPv4 адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "ip6_addr",
			translation: "Поле {0} должно быть распознаваемым IPv6 адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "unix_addr",
			translation: "Поле {0} должно быть распознаваемым UNIX адресом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "mac",
			translation: "Поле {0} должно содержать MAC адрес",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "unique",
			translation: "Поле {0} должно содержать уникальные значения",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "iscolor",
			translation: "Поле {0} должно быть цветом",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				fld, _ := ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}
				t, err := ut.T(fe.Tag(), fld)
				if err != nil {
					return fe.(error).Error()
				}
				return t
			},
		},
		{
			tag:         "oneof",
			translation: "Поле {0} должно быть одним из [{1}]",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var fld string
				fld, _ = ut.T(fe.Field())
				if fld == "" {
					fld = fe.Field()
				}

				s, err := ut.T(fe.Tag(), fld, fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "dateInFuture",
			translation: "Дата и время не могут быть в прошлом",
			override:    false,
		},
		{
			tag:         "existedEventsParams",
			translation: "Недопустимые параметры события",
			override:    false,
		},
		{
			tag:         "fileAccessType",
			translation: "Недопустимый тип доступа к файлу",
			override:    false,
		},
		{
			tag:         "starRating",
			translation: "Недопустимая оценка",
			override:    false,
		},
		{
			tag:         "userExistsInLdap",
			translation: "Пользователь не найден в LDAP",
			override:    false,
		},
		{
			tag:         "Title",
			translation: "Название",
			override:    false,
		},
		{
			tag:         "Description",
			translation: "Описание",
			override:    false,
		},
		{
			tag:         "FileName",
			translation: "Имя файла",
			override:    false,
		},
		{
			tag:         "FileAccess",
			translation: "Доступ файла",
			override:    false,
		},
		{
			tag:         "StartAt",
			translation: "Время начала",
			override:    false,
		},
		{
			tag:         "EndAt",
			translation: "Время окончания",
			override:    false,
		},
		{
			tag:         "EventID",
			translation: "ID события",
			override:    false,
		},
		{
			tag:         "UserID",
			translation: "ID пользователя",
			override:    false,
		},
		{
			tag:         "Message",
			translation: "Сообщение",
			override:    false,
		},
		{
			tag:         "Page",
			translation: "Страница",
			override:    false,
		},
		{
			tag:         "Limit",
			translation: "Лимит",
			override:    false,
		},
		{
			tag:         "Source",
			translation: "Источник",
			override:    false,
		},
		{
			tag:         "ModeratorEmails",
			translation: "Модераторы",
			override:    false,
		},
		{
			tag:         "Params",
			translation: "Параметры",
			override:    false,
		},
		{
			tag:         "Device",
			translation: "Устройство",
			override:    false,
		},
		{
			tag:         "Os",
			translation: "Операционная система",
			override:    false,
		},
		{
			tag:         "Browser",
			translation: "Браузер",
			override:    false,
		},
		{
			tag:         "Status",
			translation: "Статус",
			override:    false,
		},
		{
			tag:         "Link",
			translation: "Ссылка",
			override:    false,
		},
		{
			tag:         "ScheduleID",
			translation: "ID расписания",
			override:    false,
		},
		{
			tag:         "City",
			translation: "Город",
			override:    false,
		},
		{
			tag:         "Rating",
			translation: "Оценка",
			override:    false,
		},
		{
			tag:         "Type",
			translation: "Тип",
			override:    false,
		},
		{
			tag:         "Login",
			translation: "Логин",
			override:    false,
		},
		{
			tag:         "Password",
			translation: "Пароль",
			override:    false,
		},
		{
			tag:         "ParticipantsEmails",
			translation: "Спикеры",
			override:    false,
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)

		} else if t.customTransFunc != nil && t.customRegisFunc == nil {

			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)

		} else if t.customTransFunc == nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)

		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return

	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
