package repositories

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

func getTestProduct() entities.Product {
	id, _ := valueobjects.NewId(int64(rand.Intn(1000000-1) + 1))
	categoryName, _ := valueobjects.NewName("Category#" + uuid.NewString())
	product, _ := entities.NewProduct(int64(rand.Intn(1000000-1)+1), "Product"+uuid.NewString(), decimal.New(100, 0), "Description#123456789"+uuid.NewString(), &entities.Category{
		Id:   *id,
		Name: *categoryName,
	})
	return *product
}

type errorSessionRepository struct {
}

func NewErrorSessionRepository() SessionRepository {
	return &errorSessionRepository{}
}

func (esr *errorSessionRepository) AddSession(entities.Session) (entities.Session, error) {
	return entities.Session{}, errors.New("some error")
}

func (esr *errorSessionRepository) DeleteSession(session entities.Session) error {
	return errors.New("some error")
}

func (esr *errorSessionRepository) DeleteAllUsersSessions(userId valueobjects.Id) error {
	return errors.New("some error")
}

func (esr *errorSessionRepository) UpdateSession(session entities.Session) error {
	return errors.New("some error")
}

func (esr *errorSessionRepository) GetSession(sessionId uuid.UUID) (*entities.Session, error) {
	return nil, errors.New("some error")
}

func (esr *errorSessionRepository) GetSessionsByUserId(userId valueobjects.Id) ([]entities.Session, error) {
	return make([]entities.Session, 0), errors.New("some error")
}

func (esr *errorSessionRepository) DeleteSessionsExpiredAfter(timeDuration time.Duration) error {
	return errors.New("some error")
}

type errorUserRepository struct {
}

func NewErrorUserRepository() UserRepository {
	return &errorUserRepository{}
}

func (eur *errorUserRepository) Add(*entities.User) error {
	return errors.New("some error")
}

func (eur *errorUserRepository) Update(entities.User) error {
	return errors.New("some error")
}
func (eur *errorUserRepository) Delete(entities.User) error {
	return errors.New("some error")
}
func (eur *errorUserRepository) Get(int64) (*entities.User, error) {
	return nil, errors.New("some error")
}
func (eur *errorUserRepository) GetAll() ([]entities.User, error) {
	return make([]entities.User, 0), errors.New("some error")
}
func (eur *errorUserRepository) GetByEmail(string) (*entities.User, error) {
	return nil, errors.New("some error")
}
func (eur *errorUserRepository) ExistsByEmail(string) (bool, error) {
	return false, errors.New("some error")
}
