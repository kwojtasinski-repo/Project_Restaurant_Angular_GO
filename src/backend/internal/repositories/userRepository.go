package repositories

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type UserRepository interface {
	Add(*entities.User) error
	Update(entities.User) error
	Delete(entities.User) error
	Get(int64) (*entities.User, error)
	GetAll() ([]entities.User, error)
	GetByEmail(string) (*entities.User, error)
}

type inMemoryUserRepository struct {
	users []entities.User
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make([]entities.User, 0),
	}
}

func (repo *inMemoryUserRepository) Add(user *entities.User) error {
	var length int = len(repo.users)
	if length == 0 {
		newId, _ := valueobjects.NewId(1)
		user.Id = *newId
		repo.users = append(repo.users, *user)
		return nil
	}

	lastElement := repo.users[length-1]
	newId, _ := valueobjects.NewId(lastElement.Id.Value() + 1)
	user.Id = *newId
	repo.users = append(repo.users, *user)
	return nil
}

func (repo *inMemoryUserRepository) Update(userToUpdate entities.User) error {
	for index, user := range repo.users {
		if user.Id.Value() == userToUpdate.Id.Value() {
			user.Email = userToUpdate.Email
			user.Password = userToUpdate.Password
			user.Role = userToUpdate.Role
			repo.users[index] = user
		}
	}
	return nil
}

func (repo *inMemoryUserRepository) Delete(userToDelete entities.User) error {
	for index, user := range repo.users {
		if user.Id.Value() == userToDelete.Id.Value() {
			repo.users = append(repo.users[:index], repo.users[index+1:]...)
			return nil
		}
	}

	return nil
}

func (repo *inMemoryUserRepository) Get(id int64) (*entities.User, error) {
	for _, user := range repo.users {
		if user.Id.Value() == id {
			return &user, nil
		}
	}

	return nil, nil
}

func (repo *inMemoryUserRepository) GetAll() ([]entities.User, error) {
	return repo.users, nil
}

func (repo *inMemoryUserRepository) GetByEmail(email string) (*entities.User, error) {
	for _, user := range repo.users {
		if user.Email.Value() == email {
			return &user, nil
		}
	}

	return nil, nil
}
