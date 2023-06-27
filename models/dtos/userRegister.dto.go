package dtos

type UserRegisterDto struct {
	Lastname        string `json:"lastname" validate:"required"`
	Firstname       string `json:"firstname" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo"`
}
