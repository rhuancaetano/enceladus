package updateuser

import (
	"regexp"

	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/internal/user/dtos"
)

type UpdateUserUseCase struct {
	repo dtos.Repo
}

func NewUpdateUserUseCase(r dtos.Repo) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: r,
	}
}

func (uc *UpdateUserUseCase) execute(data *dtos.UpdateUserDTO) (*dtos.UserDTO, *usecase.UseCaseError) {
	idReg := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
	if data.Id == "" || !idReg.MatchString(data.Id) {
		return nil, usecase.BadRequestError("invalid user id")
	}

	userById, findByIdError := uc.repo.FindById(data.Id)
	if findByIdError != nil {
		return nil, usecase.ServerError()
	} else if userById == nil {
		return nil, usecase.BadRequestError("could not find user")
	}

	if data.FirstName == "" {
		return nil, usecase.BadRequestError("invalid user first name")
	}
	if data.LastName == "" {
		return nil, usecase.BadRequestError("invalid user last name")
	}
	emailReg := regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")
	if data.Email == "" || !emailReg.MatchString(data.Email) {
		return nil, usecase.BadRequestError("invalid user email")
	}

	if userById.Email != data.Email {
		userByEmail, findByEmailError := uc.repo.FindByEmail(data.Email)
		if findByEmailError != nil {
			return nil, usecase.ServerError()
		} else if userByEmail != nil {
			return nil, usecase.BadRequestError("the email is already taken")
		}
	}

	user, err := uc.repo.UpdateUser(data)
	if err != nil {
		return nil, usecase.ServerError()
	}

	return user, nil
}
