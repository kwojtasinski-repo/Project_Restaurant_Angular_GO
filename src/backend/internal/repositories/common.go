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

type errorOrderRepository struct {
}

func NewErrorOrderRepository() OrderRepository {
	return &errorOrderRepository{}
}

func (eor *errorOrderRepository) Add(*entities.Order) error {
	return errors.New("some error")
}
func (eor *errorOrderRepository) Delete(entities.Order) error {
	return errors.New("some error")
}
func (eor *errorOrderRepository) Get(int64) (*entities.Order, error) {
	return nil, errors.New("some error")
}
func (eor *errorOrderRepository) GetAllByUser(int64) ([]entities.Order, error) {
	return make([]entities.Order, 0), errors.New("some error")
}
func (eor *errorOrderRepository) GetAll() ([]entities.Order, error) {
	return make([]entities.Order, 0), errors.New("some error")
}
func (eor *errorOrderRepository) Update(*entities.Order) error {
	return errors.New("some error")
}

type errorProductRepository struct {
}

func NewErrorProductRepository() ProductRepository {
	return &errorProductRepository{}
}

func (epr *errorProductRepository) Add(*entities.Product) error {
	return errors.New("some error")
}
func (epr *errorProductRepository) Update(entities.Product) error {
	return errors.New("some error")
}
func (epr *errorProductRepository) Delete(entities.Product) error {
	return errors.New("some error")
}
func (epr *errorProductRepository) Get(int64) (*entities.Product, error) {
	return nil, errors.New("some error")
}
func (epr *errorProductRepository) GetAll() ([]entities.Product, error) {
	return make([]entities.Product, 0), errors.New("some error")
}

type errorCartRepository struct {
}

func NewErrorCartRepository() CartRepository {
	return &errorCartRepository{}
}

func (ecr *errorCartRepository) Add(cart *entities.Cart) error {
	return errors.New("some error")
}

func (ecr *errorCartRepository) Delete(cart entities.Cart) error {
	return errors.New("some error")
}

func (ecr *errorCartRepository) Get(cartId int64) (*entities.Cart, error) {
	return nil, errors.New("some error")
}

func (ecr *errorCartRepository) GetAllByUser(userId int64) ([]entities.Cart, error) {
	return make([]entities.Cart, 0), errors.New("some error")
}

func (ecr *errorCartRepository) DeleteCartByUserId(userId int64) error {
	return errors.New("some error")
}

type errorCategoryRepository struct {
}

func NewErrorCategoryRepository() CategoryRepository {
	return &errorCategoryRepository{}
}

func (ecr *errorCategoryRepository) Add(*entities.Category) error {
	return errors.New("some error")
}
func (ecr *errorCategoryRepository) Update(entities.Category) error {
	return errors.New("some error")
}
func (ecr *errorCategoryRepository) Delete(entities.Category) error {
	return errors.New("some error")
}
func (ecr *errorCategoryRepository) Get(int64) (*entities.Category, error) {
	return nil, errors.New("some error")
}
func (ecr *errorCategoryRepository) GetAll() ([]entities.Category, error) {
	return make([]entities.Category, 0), errors.New("some error")
}
