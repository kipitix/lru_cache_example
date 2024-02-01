package services

import (
	"lrucache/application/storage"
	"lrucache/domain/repository"
)

type Services interface {
	Storage() storage.Storage
	UserRepository() repository.UserRepository
}
