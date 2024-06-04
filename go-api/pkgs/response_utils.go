package pkgs

import (
	"go-api/internal/constants"
	"go-api/internal/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constants.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseKey(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Key:     status,
		Message: message,
		Data:    data,
	}
}
