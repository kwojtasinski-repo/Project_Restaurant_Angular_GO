package services

import "golang.org/x/crypto/bcrypt"

type PasswordHasherService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type passwordHasherService struct {
}

func CreatePassworHasherService() PasswordHasherService {
	return &passwordHasherService{}
}

func (hasher *passwordHasherService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 64)
	return string(bytes), err
}

func (hasher *passwordHasherService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
