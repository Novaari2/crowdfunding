package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailibility(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// buat object dari struct User sebelum di mapping
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil{
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil{
		return newUser, err
	}

	return newUser, nil
}

func (s *service)Login(input LoginInput)(User, error){
	email := input.Email
	password := input.Password

	// cari user berdasarkan login
	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	// mencocokkan password jika lolos validasi
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		return user, err
	}

	return user, nil
}

func (s *service)IsEmailAvailibility(input CheckEmailInput)(bool, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil{
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	// nilai defaultnya false
	return false, nil
}


// mapping struct input ke struct User
// simpan struct User melalui repository