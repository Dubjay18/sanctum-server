package pkg

import (
	"github.com/Dubjay18/sanctum-server/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus string, responseMessage string, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus, responseMessage, data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
