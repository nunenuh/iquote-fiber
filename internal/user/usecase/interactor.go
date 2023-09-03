package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/shared/exception"
	"github.com/nunenuh/iquote-fiber/internal/shared/hash"
	"github.com/nunenuh/iquote-fiber/internal/shared/param"
	"github.com/nunenuh/iquote-fiber/internal/shared/validator"
	"github.com/nunenuh/iquote-fiber/internal/user/domain"
)

type UserUsecase struct {
	repo      domain.IUserRepository
	validator *validator.Validator
}

func NewUserUsecase(r domain.IUserRepository) *UserUsecase {
	validator := validator.NewValidator()
	return &UserUsecase{
		repo:      r,
		validator: validator,
	}
}

func (ucase *UserUsecase) GetAll(param *param.Param) ([]*domain.User, error) {
	u, err := ucase.repo.GetAll(param.Limit, param.Page)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *UserUsecase) GetByID(ID int) (*domain.User, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *UserUsecase) GetByUsername(username string) (*domain.User, error) {
	u, err := ucase.repo.GetByUsername(username)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *UserUsecase) GetByEmail(email string) (*domain.User, error) {
	u, err := ucase.repo.GetByEmail(email)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *UserUsecase) Create(user *domain.User) (*domain.User, error) {
	err := ucase.validator.Validate(user)
	if err != nil {
		return nil, exception.NewValidatorError(err.Error())
	}

	hashedPass, err := hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPass

	u, err := ucase.repo.Create(user)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}
	return u, nil
}

func (ucase *UserUsecase) Update(ID int, user *domain.User) (*domain.User, error) {
	err := ucase.validator.Validate(user)
	if err != nil {
		return nil, exception.NewValidatorError(err.Error())
	}

	hashedPass, err := hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPass

	u, err := ucase.repo.Update(ID, user)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *UserUsecase) Delete(ID int) error {
	_, errGet := ucase.repo.GetByID(ID)
	if errGet != nil {
		return exception.NewRepositoryError(errGet.Error())
	}

	err := ucase.repo.Delete(ID)
	if err != nil {
		return exception.NewRepositoryError(err.Error())
	}

	return nil
}
