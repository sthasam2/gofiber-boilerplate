package controllers

import (
	"github.com/gofiber/fiber/v2"

	"app/serializers"
)

// HTTPResponse normalize HTTP Response format
func HTTPResponse(httpCode int, message string, data interface{}) *serializers.Response {
	return &serializers.Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}

// HTTPFiberErrorResponse normalizes fiber error responses
func HTTPFiberErrorResponse(errorObj []*fiber.Error) *serializers.FiberErrorResponse {
	// Convert fiber.Error to ErrorBody
	// This fixes issues with swagger auto generated docs not identify fiber.Error type
	var errorSlice []*serializers.ErrorBody
	for i := 0; i < len(errorObj); i++ {
		errorSlice = append(errorSlice, mapToErrorOutput(errorObj[i]))
	}

	return &serializers.FiberErrorResponse{
		Error: errorSlice,
	}
}

// HTTPErrorResponse normalizes error responses
func HTTPErrorResponse(errorObj []*serializers.Response) *serializers.ErrorResponse {
	var errorSlice []*serializers.Response
	for i := 0; i < len(errorObj); i++ {
		errorSlice = append(errorSlice, errorObj[i])
	}

	return &serializers.ErrorResponse{
		Error: errorSlice,
	}
}

// SendErrorHTTPResponse sends Error HTTP response
func SendErrorHTTPResponse(c *fiber.Ctx, httpCode int, message string, err interface{}) error {

	// checking err interface using type assertions

	// Response type
	if errData, ok := err.([]*serializers.Response); ok {
		// errBody
		return c.Status(httpCode).JSON(HTTPErrorResponse(errData))
	}

	// Fiber Error
	if errData, ok := err.([]*fiber.Error); ok {
		return c.Status(httpCode).JSON(HTTPFiberErrorResponse(errData))
	}

	// Fallback Default
	var errData = &serializers.Response{Code: httpCode, Message: message, Data: err}
	return c.Status(httpCode).JSON(HTTPErrorResponse([]*serializers.Response{errData}))
}

func SendSuccessHTTPResponse(c *fiber.Ctx, httpCode int, message string, data interface{}) error {
	return c.Status(httpCode).JSON(HTTPResponse(httpCode, message, data))
}

//////////////////////
//
//////////////////////

//////////////////////
// Private Method
//////////////////////

func mapToErrorOutput(err *fiber.Error) *serializers.ErrorBody {
	return &serializers.ErrorBody{
		Code:    err.Code,
		Message: err.Message,
	}
}
