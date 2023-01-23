package serializers

import (
	"app/models"

	"github.com/google/uuid"
)

type UserOutput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

// UserRegisterInput defines data structure for input
type UserRegisterInput struct {
	Username string `json:"username" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,min=5,max=100,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginInput struct {
	Email    string `json:"email" validate:"required,min=5,max=100,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginOutput struct {
	Access  string     `json:"access"`
	Refresh string     `json:"refresh"`
	User    UserOutput `json:"user"`
}

// type UserUpdate struct {
// }

// type UserDelete struct {
// }

// type UserToggleActive struct {
// }

// type UserResetPassword struct {
// }

//////////////////////////////////////////////////
// Mapping Methods i.e. Serialize data to models
//////////////////////////////////////////////////

func MapRegisterInputToUser(userInput UserRegisterInput) models.User {
	return models.User{
		Username:   userInput.Username,
		Email:      userInput.Email,
		Password:   userInput.Password,
		ExternalID: uuid.New().String(),
	}
}

func MapUserToOutPut(u *models.User) *UserOutput {
	return &UserOutput{
		ID:       u.ExternalID,
		Email:    u.Email,
		Username: u.Username,
	}
}

func MapUserAndTokensToOutput(u *models.User, aT TokenDetails, rT TokenDetails) {
	// return

}
