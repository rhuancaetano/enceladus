package createuser

import (
	"regexp"

	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/internal/user/dtos"
)

type CreateUserUseCase struct {
	repo dtos.Repo
}

func NewCreateUserUseCase(r dtos.Repo) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: r,
	}
}

func (uc *CreateUserUseCase) execute(data *dtos.CreateUserDTO) (*dtos.UserDTO, *usecase.UseCaseError) {
	if data.FirstName == "" {
		return nil, usecase.BadRequestError("invalid user first name")
	}
	if data.LastName == "" {
		return nil, usecase.BadRequestError("invalid user last name")
	}
	reg := regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")
	if data.Email == "" || !reg.MatchString(data.Email) {
		return nil, usecase.BadRequestError("invalid user email")
	}
	if user, err := uc.repo.FindByEmail(data.Email); err != nil {
		return nil, usecase.ServerError()
	} else if user != nil {
		return nil, usecase.BadRequestError("the email is already taken")
	}

	user, err := uc.repo.CreateUser(data)
	if err != nil {
		return nil, usecase.ServerError()
	}

	return user, nil
}
