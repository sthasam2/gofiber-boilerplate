package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Validate validates the parsed body for errors
func Validate(payload interface{}) []*fiber.Error {

	// validate a struct
	if err := validate.Struct(payload); err != nil {
		var errorList []*fiber.Error

		// typecasting ie err.() for getting error as suggested by documentation
		for _, err := range err.(validator.ValidationErrors) {
			errorList = append(errorList,
				&fiber.Error{
					Code:    fiber.StatusBadRequest,
					Message: fmt.Sprintf("%v must be valid", err.StructField()),
				})
		}

		return errorList
	}

	return nil
}

// ParseBody parses request body from context and maps it to body interface struct
func ParseBody(c *fiber.Ctx, body interface{}) []*fiber.Error {

	// bind request body to a struct and check errors
	if err := c.BodyParser(body); err != nil {
		var errorList []*fiber.Error
		errorList = append(
			errorList,
			fiber.ErrBadRequest,
		)

		return errorList
	}

	return nil
}

// ParseBodyandValidate applies ParseBudy and then Validate
func ParseBodyandValidate(c *fiber.Ctx, body interface{}) []*fiber.Error {

	if err := ParseBody(c, body); err != nil {
		return err
	}

	return Validate(body)

}
