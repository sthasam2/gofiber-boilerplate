package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"app/serializers"
	"app/services"
	"app/utils"
)

// RegisterUser Godoc
// @Summary Register User
// @Description Registers a user
// @Tags Auth
// @Produce json
// @Param payload body serializers.UserRegisterInput true "Register Body"
// @Success 201 {object} serializers.Response
// @Failure 400 {array} serializers.ErrorResponse
// @Failure 401 {array} serializers.ErrorResponse
// @Failure 500 {array} serializers.ErrorResponse
// @Router /api/auth/register [post]
func RegisterUser(c *fiber.Ctx) error {
	var userRegisterInput serializers.UserRegisterInput

	// parse and validate request body
	if err := utils.ParseBodyandValidate(c, &userRegisterInput); err != nil {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusBadRequest,
			"Request Body Error",
			HTTPFiberErrorResponse(err),
		)
	}

	// check user already exists
	emailUserExists := CheckUserExistsByEmail(userRegisterInput.Email)
	usernameUserExists := CheckUserExistsByUsername(userRegisterInput.Username)

	if emailUserExists || usernameUserExists {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusBadRequest,
			"Unable to register User",
			fmt.Sprintf(
				"User with credentails already exists. "+
					"Email '%s' available: %t. "+
					"Useraname '%s' available: %t",
				userRegisterInput.Email, !emailUserExists,
				userRegisterInput.Username, !usernameUserExists,
			),
		)
	}

	// map to model
	newUser := serializers.MapRegisterInputToUser(userRegisterInput)

	// hash password
	hashedPW, _ := utils.HashPassword(userRegisterInput.Password)
	newUser.Password = hashedPW

	// save to DB
	if err := CreateUser(&newUser); err != nil {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusInternalServerError,
			"Unable to register User",
			err,
		)
	}

	// TODO
	// Create token and send emails

	// if not error newUser has been created
	return SendSuccessHTTPResponse(
		c,
		fiber.StatusCreated,
		"User created",
		serializers.MapUserToOutPut(&newUser),
	)

}

// LoginUser Godoc
// @Summary User Login
// @Description Logs in a user
// @Tags Auth
// @Produce json
// @Param payload body serializers.UserRegisterInput true "Login Body"
// @Success 201 {object} serializers.Response
// @Failure 400 {array} serializers.ErrorResponse
// @Failure 401 {array} serializers.ErrorResponse
// @Failure 500 {array} serializers.ErrorResponse
// @Router /api/auth/login [post]
func LoginController(c *fiber.Ctx) error {
	var userLoginInput serializers.UserLoginInput

	// Parse and Validate
	if err := utils.ParseBodyandValidate(c, &userLoginInput); err != nil {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusBadRequest,
			"Could not Login",
			HTTPFiberErrorResponse(err),
		)
	}

	// Get user using input email
	user, err := GetUserByEmail(userLoginInput.Email)
	if err != nil {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusNotFound,
			fmt.Sprintf("User with email '%s' not found", userLoginInput.Email),
			err,
		)
	}

	// Check password
	if passwordIsCorrect := utils.CheckPasswordHash(userLoginInput.Password, user.Password); !passwordIsCorrect {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusUnauthorized,
			"Email or Password is Incorrect",
			nil,
		)
	}

	// Tokens
	access, err1 := services.IssueAccessToken(*user)
	refresh, err2 := services.IssueRefreshToken(*user)

	if err1 != nil || err2 != nil {
		return SendErrorHTTPResponse(
			c,
			fiber.StatusInternalServerError,
			"Something went wrong: Could not create token",
			err,
		)
	}

	return SendSuccessHTTPResponse(
		c,
		fiber.StatusOK,
		"Login Success",
		fiber.Map{
			"user":    serializers.MapUserToOutPut(user),
			"access":  access.Token,
			"refersh": refresh.Token,
		})
}
