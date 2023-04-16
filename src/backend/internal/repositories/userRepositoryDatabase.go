package repositories

import (
	"database/sql"
	"errors"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type userRepository struct {
	database sql.DB
}

func CreateUserRepository(database sql.DB) UserRepository {
	return &userRepository{
		database: database,
	}
}

func (repo *userRepository) Add(user *entities.User) error {
	query := "INSERT INTO `users` (email, password, role, deleted) VALUES (?, ?, ?, ?);"
	_, err := repo.database.Exec(query, user.Email.Value(), user.Password, user.Role, user.Deleted)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Update(userToUpdate entities.User) error {
	query := "UPDATE `users` SET password = ?, role = ?, deleted = ?;"
	_, err := repo.database.Exec(query, userToUpdate.Password, userToUpdate.Role, userToUpdate.Deleted)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Delete(userToDelete entities.User) error {
	query := "DELETE FROM `users` WHERE id = ?;"
	_, err := repo.database.Exec(query, userToDelete.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Get(id int64) (*entities.User, error) {
	query := "SELECT id, email, password, role, deleted FROM `users` WHERE id = ?;"
	row := repo.database.QueryRow(query, id)
	user, err := getUser(row)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) GetAll() ([]entities.User, error) {
	users := make([]entities.User, 0)
	query := "SELECT id, email, password, role FROM `users` WHERE deleted = false;"
	rows, err := repo.database.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId int64
		var email string
		var password string
		var role string
		if err := rows.Scan(&userId, &email, &password, &role); err != nil {
			return nil, err
		}

		user, _ := entities.NewUser(userId, email, password, role)
		users = append(users, *user)
	}

	return users, nil
}

func (repo *userRepository) GetByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, password, role, deleted FROM `users` WHERE email = ?;"
	row := repo.database.QueryRow(query, email)
	user, err := getUser(row)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) ExistsByEmail(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM `users` WHERE email = ?);"
	var exists bool = false
	row := repo.database.QueryRow(query, email)
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func getUser(row *sql.Row) (*entities.User, error) {
	if row == nil {
		return nil, errors.New("getting User -> Row is nil")
	}

	var userId int64
	var emailFromDb string
	var password string
	var role string
	var deleted bool
	if err := row.Scan(&userId, &emailFromDb, &password, &role, &deleted); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	newUserId, _ := valueobjects.NewId(userId)
	newEmail, _ := valueobjects.NewEmailAddress(emailFromDb)
	user := &entities.User{
		Id:       *newUserId,
		Email:    *newEmail,
		Password: password,
		Role:     role,
		Deleted:  deleted,
	}
	return user, nil
}
