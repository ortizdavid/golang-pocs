package container

import (
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/repositories"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	UserRepository    *repositories.UserRepository
	ProductRepository *repositories.ProductRepository
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		UserRepository:    repositories.NewUserRepository(db),
		ProductRepository: repositories.NewProductRepository(db),
	}
}