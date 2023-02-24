package repositoryuser

import (
	"errors"

	entityuser "github.com/faqihyugos/penaduan-api/entities/entityUser"
	"gorm.io/gorm"
)

type RepositoryUser interface {
	Create(data entityuser.User) (entityuser.User, error)
	IsEmailExist(email string) error
	Login(email string) (entityuser.User, error)
	Update(data entityuser.User) (entityuser.User, error)
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

// DeleteByID implements RepositoryUser
func (*repository) DeleteByID(id uint) error {
	panic("unimplemented")
}

// Login implements RepositoryUser
func (*repository) Login(email string) (entityuser.User, error) {
	panic("unimplemented")
}

// Update implements RepositoryUser
func (*repository) Update(data entityuser.User) (entityuser.User, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) RepositoryUser {
	return &repository{db: db}
}

func (r *repository) Create(data entityuser.User) (entityuser.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entityuser.User{}, err
	}

	return data, nil
}

func (r *repository) IsEmailExist(email string) error {
	user := new(entityuser.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}
	return errors.New("email already exist")
}
