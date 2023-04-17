package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
)

type SessionService interface {
	CreateSession(user entities.User) (*dto.SessionDto, *applicationerrors.ErrorStatus)
	RevokeSession(sessionId uuid.UUID) *applicationerrors.ErrorStatus
	RevokeAllUsersSessions(userId int64) *applicationerrors.ErrorStatus
	RefreshSession(sessionId uuid.UUID) (*dto.SessionDto, *applicationerrors.ErrorStatus)
	GetUserSessions(userId int64) ([]dto.SessionDto, *applicationerrors.ErrorStatus)
	ManageSession(sessionDto dto.SessionDto) (*dto.SessionDto, *applicationerrors.ErrorStatus)
	ClearExpiredSessions() *applicationerrors.ErrorStatus
}

type sessionService struct {
	repo     repositories.SessionRepository
	userRepo repositories.UserRepository
}

func CreateSessionService(sessionRepository repositories.SessionRepository, userRepository repositories.UserRepository) SessionService {
	return &sessionService{
		repo:     sessionRepository,
		userRepo: userRepository,
	}
}

func (service *sessionService) CreateSession(user entities.User) (*dto.SessionDto, *applicationerrors.ErrorStatus) {
	session := entities.CreateSession(user, createTokenLifetime())
	var err error
	session, err = service.repo.AddSession(session)

	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	sessionDto := dto.MapToSessionDto(session)
	return &sessionDto, nil
}

func (service *sessionService) RevokeSession(sessionId uuid.UUID) *applicationerrors.ErrorStatus {
	var session *entities.Session
	var err error

	session, err = service.repo.GetSession(sessionId)

	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	if session == nil {
		return applicationerrors.UnAuthorized()
	}

	service.repo.DeleteSession(*session)
	return nil
}

func (service *sessionService) RefreshSession(sessionId uuid.UUID) (*dto.SessionDto, *applicationerrors.ErrorStatus) {
	session, err := service.repo.GetSession(sessionId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if session == nil {
		return nil, applicationerrors.UnAuthorized()
	}

	userId := session.UserId()
	user, errUserRepo := service.userRepo.Get(userId.Value())
	if errUserRepo != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	session.SetUser(*user)
	session.SetExpiry(createTokenLifetime())
	service.repo.UpdateSession(*session)

	sessionDto := dto.MapToSessionDto(*session)
	return &sessionDto, nil
}

func (service *sessionService) GetUserSessions(userId int64) ([]dto.SessionDto, *applicationerrors.ErrorStatus) {
	var newUserId *valueobjects.Id
	var err error
	var sessions []entities.Session
	newUserId, err = valueobjects.NewId(userId)

	if err != nil {
		return nil, applicationerrors.BadRequest("Invalid UserId")
	}

	sessions, err = service.repo.GetSessionsByUserId(*newUserId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	sessionsDto := make([]dto.SessionDto, 0)
	for _, session := range sessions {
		sessionsDto = append(sessionsDto, dto.MapToSessionDto(session))
	}

	return sessionsDto, nil
}

func (service *sessionService) ManageSession(sessionDto dto.SessionDto) (*dto.SessionDto, *applicationerrors.ErrorStatus) {
	session, err := service.repo.GetSession(sessionDto.SessionId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}
	if session == nil {
		return nil, applicationerrors.UnAuthorizedWithMessage("Invalid Cookie")
	}

	userId := session.UserId()
	if userId.Value() != sessionDto.UserId {
		return nil, applicationerrors.UnAuthorized()
	}

	exists, errRepo := service.userRepo.ExistsByEmail(sessionDto.Email)
	if errRepo != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if !exists {
		return nil, applicationerrors.UnAuthorized()
	}

	sessionTime := time.UnixMilli(sessionDto.Expiry)
	if sessionTime.After(time.Now().UTC()) {
		sessionDto := dto.MapToSessionDto(*session)
		return &sessionDto, nil
	}

	refreshedSessionDto, errStatus := service.RefreshSession(sessionDto.SessionId)
	if errStatus != nil {
		return nil, errStatus
	}

	return refreshedSessionDto, nil
}

func (service *sessionService) RevokeAllUsersSessions(userId int64) *applicationerrors.ErrorStatus {
	newUserId, err := valueobjects.NewId(userId)
	if err != nil {
		applicationerrors.BadRequest("Invalid UserId")
	}

	err = service.repo.DeleteAllUsersSessions(*newUserId)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func (service *sessionService) ClearExpiredSessions() *applicationerrors.ErrorStatus {
	if err := service.repo.DeleteSessionsExpiredAfter(time.Duration(settings.CookieLifeTime)); err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func createTokenLifetime() time.Time {
	return time.Now().UTC().Add(time.Hour * 2) // 2 hours
}
