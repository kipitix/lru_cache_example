package services

import (
	"lrucache/application/storage"
	"lrucache/domain/repository"
)

type servicesImpl struct {
	storage  storage.Storage
	userRepo repository.UserRepository
}

func New(storage storage.Storage, userRepo repository.UserRepository) Services {
	return &servicesImpl{
		storage:  storage,
		userRepo: userRepo,
	}
}

var _ Services = (*servicesImpl)(nil)

func (s servicesImpl) Storage() storage.Storage {
	return s.storage
}

func (s servicesImpl) UserRepository() repository.UserRepository {
	return s.userRepo
}
