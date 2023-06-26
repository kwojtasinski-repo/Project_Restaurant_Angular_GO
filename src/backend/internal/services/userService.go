package services

import (
	"fmt"
	"strings"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
)

type UserService interface {
	Register(*dto.AddUserDto) (*dto.UserDto, *applicationerrors.ErrorStatus)
	Delete(userId int64) *applicationerrors.ErrorStatus
	Get(userId int64) (*dto.UserDto, *applicationerrors.ErrorStatus)
	GetAll() ([]dto.UserDto, *applicationerrors.ErrorStatus)
	Login(dto.SignInDto) (*dto.SessionDto, *applicationerrors.ErrorStatus)
}

type userService struct {
	repository     repositories.UserRepository
	passwordHasher PasswordHasherService
	sessionService SessionService
}

func CreateUserService(repo repositories.UserRepository, passwordHasher PasswordHasherService, sessionService SessionService) UserService {
	return &userService{
		repository:     repo,
		passwordHasher: passwordHasher,
		sessionService: sessionService,
	}
}

func (service *userService) Register(addUser *dto.AddUserDto) (*dto.UserDto, *applicationerrors.ErrorStatus) {
	if addUser == nil {
		return nil, applicationerrors.BadRequest(applicationerrors.InvalidUser)
	}

	if err := addUser.Validate(); err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	addUser.Email = strings.ToLower(addUser.Email)
	exists, errRepo := service.repository.ExistsByEmail(addUser.Email)
	if errRepo != nil {
		return nil, applicationerrors.InternalError(errRepo.Error())
	}

	if exists {
		return nil, applicationerrors.BadRequest(applicationerrors.InvalidEmailPassword)
	}

	hashedPassword, err := service.passwordHasher.HashPassword(addUser.Password)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	user, err := entities.NewUser(0, addUser.Email, hashedPassword, "user")
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	err = service.repository.Add(user)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	return dto.MapToUserDto(*user), nil
}

func (service *userService) Delete(id int64) *applicationerrors.ErrorStatus {
	user, err := service.repository.Get(id)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	if user == nil {
		return applicationerrors.NotFoundWithMessage(fmt.Sprintf("'User' with id %v was not found", id))
	}

	if errService := service.sessionService.RevokeAllUsersSessions(user.Id.Value()); errService != nil {
		return errService
	}

	err = service.repository.Delete(*user)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func (service *userService) Get(id int64) (*dto.UserDto, *applicationerrors.ErrorStatus) {
	user, err := service.repository.Get(id)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if user == nil {
		return nil, applicationerrors.NotFound()
	}

	return dto.MapToUserDto(*user), nil
}

func (service *userService) GetAll() ([]dto.UserDto, *applicationerrors.ErrorStatus) {
	users, err := service.repository.GetAll()
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	usersDto := make([]dto.UserDto, 0)
	for _, userInRepo := range users {
		usersDto = append(usersDto, dto.UserDto{
			Id:    dto.IdObject{ValueInt: userInRepo.Id.Value()},
			Email: userInRepo.Email.Value(),
			Role:  userInRepo.Role,
		})
	}

	return usersDto, nil
}

func (service *userService) Login(signInDto dto.SignInDto) (*dto.SessionDto, *applicationerrors.ErrorStatus) {
	user, err := service.repository.GetByEmail(strings.ToLower(signInDto.Email))
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if user == nil {
		return nil, applicationerrors.BadRequest(applicationerrors.InvalidCredentials)
	}

	matched := service.passwordHasher.CheckPasswordHash(signInDto.Password, user.Password)
	if !matched {
		return nil, applicationerrors.BadRequest(applicationerrors.InvalidCredentials)
	}

	session, errSession := service.sessionService.CreateSession(*user)
	if errSession != nil {
		return nil, errSession
	}

	return session, nil
}
