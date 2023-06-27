package mappers

import (
	"strings"

	"github.com/andy10134/asisPay-backend/models/dtos"
	"github.com/andy10134/asisPay-backend/models/entities"
	"golang.org/x/crypto/bcrypt"
)

func DtoRegisterToEntity(toMap dtos.UserRegisterDto) (entities.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(toMap.Password), bcrypt.DefaultCost)

	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		Firstname: toMap.Firstname,
		Lastname:  toMap.Lastname,
		Email:     strings.ToLower(toMap.Email),
		Password:  string(hashedPassword),
		Photo:     &toMap.Photo,
	}, nil
}

func EntityToResponse(user *entities.User) dtos.UserResponse {
	return dtos.UserResponse{
		ID:        *user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Role:      *user.Role,
		Photo:     *user.Photo,
		Provider:  *user.Provider,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}
