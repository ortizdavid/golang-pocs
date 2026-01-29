package services

import (
	"errors"
	"github.com/ortizdavid/golang-pocs/go-messagepack/models"
	"github.com/ortizdavid/golang-pocs/go-messagepack/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

type OperationResult struct {
	ID      int    `json:"id" msgpack:"id"`
	Message string `json:"message" msgpack:"message"`
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

// --- Métodos de Escrita ---

func (s *ProductService) Create(p *models.ProductModel) (OperationResult, error) {
	exists, _ := s.repository.Exists(p.Code)
	if exists {
		return OperationResult{}, errors.New("product code already exists")
	}

	if err := s.repository.Create(p); err != nil {
		return OperationResult{}, err
	}

	return OperationResult{ID: p.ID, Message: "Product successfully created"}, nil
}

func (s *ProductService) Update(id int, p *models.ProductModel) (OperationResult, error) {
	// 1. Verificação de Existência (Fail Fast)
	exists, _ := s.repository.GetByID(id)
	if exists == nil {
		return OperationResult{}, errors.New("product not found")
	}

	// 2. Executa a Atualização
	p.ID = id
	if err := s.repository.Update(*p); err != nil {
		return OperationResult{}, err
	}

	return OperationResult{ID: id, Message: "Product updated"}, nil
}

func (s *ProductService) Delete(id int) (OperationResult, error) {
	// 1. Verificação de Existência
	exists, _ := s.repository.GetByID(id)
	if exists == nil {
		return OperationResult{}, errors.New("product not found")
	}

	// 2. Executa a Remoção
	if err := s.repository.Delete(id); err != nil {
		return OperationResult{}, err
	}

	return OperationResult{ID: id, Message: "Product deleted"}, nil
}

// --- Métodos de Leitura ---

func (s *ProductService) GetAll() ([]models.ProductModel, error) {
	return s.repository.GetAll()
}

func (s *ProductService) GetByID(id int) (models.ProductModel, error) {
	p, err := s.repository.GetByID(id)
	if err != nil || p == nil {
		return models.ProductModel{}, errors.New("product not found")
	}
	return *p, nil
}