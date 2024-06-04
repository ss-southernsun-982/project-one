package pkgs

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/internal/constants"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string   `json:"message"`
	Key     string   `json:"key"`
	Errors  []string `json:"errors"`
}

func PanicException_(key string, message string, err error) {
	// err := errors.New(message)
	var errR ErrorResponse = ErrorResponse{
		Message: message,
		Key:     key,
		Errors:  []string{err.Error()},
	}
	var errType reflect.Type = reflect.TypeOf(err)
	switch errType {
	case
		reflect.TypeOf(validator.ValidationErrors{}):
		errR = validationError(err, message, key)
	case
		reflect.TypeOf(&json.SyntaxError{}):
		errR = jsonSyntaxError(err)
	case reflect.TypeOf(&json.UnmarshalTypeError{}):
		errR = unmarshalError(err, message, key)
	}

	// err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		enCodeJson, _ := json.Marshal(&errR)
		panic(enCodeJson)
	}
}

func validationError(err error, message, key string) ErrorResponse {
	var errs []string = []string{}
	var errsValidator validator.ValidationErrors
	errors.As(err, &errsValidator)
	for _, e := range errsValidator {
		errorFormat := strings.TrimSpace(fmt.Sprintf("%s is invalid, must required %s %s %s", e.Value(), e.Tag(), e.Kind(), e.Param()))
		errs = append(errs, errorFormat)
	}
	return ErrorResponse{
		Message: message,
		Key:     key,
		Errors:  errs,
	}
}

func unmarshalError(err error, message, key string) ErrorResponse {
	var errsUnmarshal *json.UnmarshalTypeError
	errors.As(err, &errsUnmarshal)
	errs := fmt.Sprintf("%s must be type %s", errsUnmarshal.Field, errsUnmarshal.Type)
	return ErrorResponse{
		Message: message,
		Key:     key,
		Errors:  []string{errs},
	}
}

func jsonSyntaxError(err error) ErrorResponse {
	var errJsonFormat *json.SyntaxError
	errors.As(err, &errJsonFormat)
	errs := fmt.Sprintf("%s", errJsonFormat)
	return ErrorResponse{
		Message: constants.InvalidRequest.GetResponseMessage(),
		Key:     constants.InvalidRequest.GetResponseKey(),
		Errors:  []string{errs},
	}
}

func PanicException(responseKey constants.ResponseStatus, err error) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage(), err)
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		errorString := string(err.([]byte))
		fmt.Println(errorString)
		var errR ErrorResponse
		oke := json.Unmarshal([]byte(errorString), &errR)
		if oke != nil {
			c.JSON(http.StatusInternalServerError, BuildResponse_(constants.InvalidRequest.GetResponseStatus(), constants.InvalidRequest.GetResponseMessage(), oke))
			c.Abort()
			return
		}
		key := errR.Key
		msg := errR.Message
		errs := errR.Errors
		switch key {
		case
			constants.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, errs))
			c.Abort()
		case
			constants.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, errs))
			c.Abort()
		case
			constants.BadRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, errs))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, errs))
			c.Abort()
		}
	}
}
