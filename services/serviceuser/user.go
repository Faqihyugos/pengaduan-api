package serviceuser

import (
	entityuser "github.com/faqihyugos/pengaduan-api/entities/entityUser"
	repositoryuser "github.com/faqihyugos/pengaduan-api/repositories/repositoryUser"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	Create(data entityuser.Request) (entityuser.Response, error)
	Login(data entityuser.RequestLogin) (entityuser.ResponseLogin, error)
	Update(data entityuser.Request) (entityuser.Response, error)
	DeleteByID(id uint) error
}

type service struct {
	repo repositoryuser.RepositoryUser
}

func New(repo repositoryuser.RepositoryUser) ServiceUser {
	return &service{repo: repo}
}

// Create implements ServiceUser
func (s *service) Create(input entityuser.Request) (entityuser.Response, error) {
	// validate input
	err := input.Validate()
	if err != nil {
		return entityuser.Response{}, err
	}

	entityUser := new(entityuser.User)
	entityUser.Fullname = input.Fullname
	entityUser.Email = input.Email
	entityUser.Role = input.Role
	entityUser.Nik = input.Nik
	entityUser.Phone = input.Phone
	entityUser.Username = input.Username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return entityuser.Response{}, err
	}
	entityUser.Password = string(hashedPassword)

	NewUser, err := s.repo.Create(*entityUser)
	if err != nil {
		return entityuser.Response{}, err
	}

	// map entity to response
	response := entityuser.Response{}
	response.ID = NewUser.ID
	response.Fullname = NewUser.Fullname
	response.Email = NewUser.Email
	response.Role = NewUser.Role
	response.Nik = NewUser.Nik
	response.Phone = NewUser.Phone
	response.Username = NewUser.Username

	return response, nil
}

// DeleteByID implements ServiceUser
func (*service) DeleteByID(id uint) error {
	panic("unimplemented")
}

// Login implements ServiceUser
func (*service) Login(data entityuser.RequestLogin) (entityuser.ResponseLogin, error) {
	panic("unimplemented")
}

// Update implements ServiceUser
func (*service) Update(data entityuser.Request) (entityuser.Response, error) {
	panic("unimplemented")
}
