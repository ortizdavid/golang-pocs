package services

import (
	"errors"
	"github.com/ortizdavid/golang-pocs/go-jsonrpc/products/models"
	"github.com/ortizdavid/golang-pocs/go-jsonrpc/products/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

type OperationResult struct {
	ID      int
	Message string
}

type StockArgs struct {
	ProductID int
	Quantity int
}

type CalculationResult struct {
    ProductName string  
    UnitPrice   float64 
    Quantity    int    
    TotalPrice  float64
}


func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

// --- Exposed methods ---
func (s *ProductService) Create(p *models.ProductModel, reply *OperationResult) error {
	exists, _ := s.repository.Exists(p.Code)
	if exists {
		return errors.New("product code already exists")
	}
	err := s.repository.Create(p)
	if err != nil {
		return err
	}
	*reply = OperationResult{
		ID:      p.ID, 
		Message: "Product successfully created",
	}
	return nil
}

func (s *ProductService) Update(p *models.ProductModel, reply *OperationResult) error {
	err := s.repository.Update(*p)
	if err != nil {
		return err
	}
	*reply = OperationResult{
		ID:      p.ID,
		Message: "Product updated",
	}
	return nil
}

func (s *ProductService) Delete(id *int, reply *OperationResult) error {
	err := s.repository.Delete(*id)
	if err != nil {
		return err
	}
	*reply = OperationResult{
		ID:      *id,
		Message: "Product deleted",
	}
	return nil
}

func (s *ProductService) GetAll(args *struct{}, reply *[]models.ProductModel) error {
    products, err := s.repository.GetAll()
    if err != nil {
        return err
    }
    *reply = products
    return nil
}

func (s *ProductService) GetByID(id *int, reply *models.ProductModel) error {
	p, err := s.repository.GetByID(*id)
	if err != nil {
		return err
	}
	*reply = *p
	return nil
}

func (s *ProductService) CalculateTotal(args *StockArgs, reply *CalculationResult) error {
    product, err := s.repository.GetByID(args.ProductID)
    if err != nil {
        return err
    }
    *reply = CalculationResult{
        ProductName: product.Name,
        UnitPrice:   product.UnitPrice,
        Quantity:    args.Quantity,
        TotalPrice:  product.UnitPrice * float64(args.Quantity),
    }
    return nil
}

