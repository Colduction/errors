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

func NewEmptyFieldErr(f interface{}) ErrEmptyField {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrEmptyField(fn)
}

func NewEmptyInputErr(f interface{}) ErrEmptyInput {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrEmptyInput(fn)
}

func NewEmptyOrInvalidFieldErr(f interface{}) ErrEmptyOrInvalidField {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrEmptyOrInvalidField(fn)
}

func NewEmptyOrInvalidInputErr(f interface{}) ErrEmptyOrInvalidInput {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrEmptyOrInvalidInput(fn)
}

func NewFieldNotFoundErr(f interface{}) ErrFieldNotFound {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrFieldNotFound(fn)
}

func NewNullInputErr(f interface{}) ErrNullInput {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrNullInput(fn)
}

func NewNullStructErr(f interface{}) ErrNullStruct {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrNullStruct(fn)
}

func NewInvalidJSONErr(f interface{}) ErrInvalidJSON {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ErrInvalidJSON(fn)
}
