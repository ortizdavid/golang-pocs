package container

import (
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/resources"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/services"
)

type ServiceContainer struct {
	UserService    *services.UserService
	ProductService *services.ProductService
}

func NewServiceContainer(repos *RepositoryContainer, res *resources.InfraResources) *ServiceContainer {
	return &ServiceContainer{
		UserService:    services.NewUserService(repos.UserRepository, res.Logger),
		ProductService: services.NewProductService(repos.ProductRepository, res.Cache, res.Logger),
	}
}