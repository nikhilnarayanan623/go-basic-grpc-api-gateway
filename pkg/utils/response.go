package utils

import (
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
	"strings"
)

func SuccessResponse(statusCode uint32, message string, data ...interface{}) domain.Response {
	return domain.Response{
		StatusCode: statusCode,
		Message:    message,
		Error:      nil,
		Data:       data,
	}
}

func ErrorResponse(statusCode uint32, message string, err string, data interface{}) domain.Response {
	return domain.Response{
		StatusCode: statusCode,
		Message:    message,
		Error:      strings.Split(err, "/n"),
	}
}
