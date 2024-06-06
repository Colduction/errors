package errors

import (
	"reflect"
	"runtime"
	"strconv"
)

type (
	ErrEmptyField            string
	ErrEmptyInput            string
	ErrEmptyOrInvalidField   string
	ErrEmptyOrInvalidInput   string
	ErrFieldNotFound         string
	ErrNullInput             string
	ErrNullStruct            string
	ErrInvalidJSON           string
	ErrInvalidJSONFieldValue string
	ErrJSONFieldNotExist     string
)

func (e ErrEmptyField) Error() string {
	return strconv.Quote(string(e)) + " field is empty"
}

func (e ErrEmptyInput) Error() string {
	return strconv.Quote(string(e)) + " function has received an empty input"
}

func (e ErrEmptyOrInvalidField) Error() string {
	return strconv.Quote(string(e)) + " field is empty or invalid"
}

func (e ErrEmptyOrInvalidInput) Error() string {
	return strconv.Quote(string(e)) + " function has received an empty or invalid input"
}

func (e ErrFieldNotFound) Error() string {
	return strconv.Quote(string(e)) + " field doesn't exist"
}

func (e ErrNullInput) Error() string {
	return strconv.Quote(string(e)) + " function has received a null input"
}

func (e ErrNullStruct) Error() string {
	return strconv.Quote(string(e)) + " struct is null"
}

func (e ErrInvalidJSON) Error() string {
	return strconv.Quote(string(e)) + " function has received an invalid json"
}

func (e ErrInvalidJSONFieldValue) Error() string {
	return strconv.Quote(string(e)) + " json field has invalid value"
}

func (e ErrJSONFieldNotExist) Error() string {
	return strconv.Quote(string(e)) + " json field is not exist"
}

func NewEmptyFieldErr(field interface{}) ErrEmptyField {
	fn := reflect.Indirect(reflect.ValueOf(field)).Type().Name()
	return ErrEmptyField(fn)
}

func NewEmptyInputErr(f interface{}) ErrEmptyInput {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrEmptyInput(fn)
}

func NewEmptyOrInvalidFieldErr(field interface{}) ErrEmptyOrInvalidField {
	fn := reflect.Indirect(reflect.ValueOf(field)).Type().Name()
	return ErrEmptyOrInvalidField(fn)
}

func NewEmptyOrInvalidInputErr(f interface{}) ErrEmptyOrInvalidInput {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrEmptyOrInvalidInput(fn)
}

func NewFieldNotFoundErr(s string) ErrFieldNotFound {
	return ErrFieldNotFound(s)
}

func NewNullInputErr(f interface{}) ErrNullInput {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrNullInput(fn)
}

func NewNullStructErr(s string) ErrNullStruct {
	return ErrNullStruct(s)
}

func NewInvalidJSONErr(f interface{}) ErrInvalidJSON {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrInvalidJSON(fn)
}

func NewInvalidJSONFieldValueErr(s string) ErrInvalidJSONFieldValue {
	return ErrInvalidJSONFieldValue(s)
}

func NewJSONFieldNotExistErr(s string) ErrJSONFieldNotExist {
	return ErrJSONFieldNotExist(s)
}
