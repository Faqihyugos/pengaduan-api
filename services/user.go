package services

import (
	entityuser "github.com/faqihyugos/penaduan-api/entities/entityUser"
	repositoryuser "github.com/faqihyugos/penaduan-api/repositories/repositoryUser"
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
