package repositories

import (
	"database/sql"

	"github.com/ortizdavid/golang-pocs/go-messagepack/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return  &ProductRepository{
		DB: db,
	}
}

func (repo *ProductRepository) Create(p *models.ProductModel) error {
    query := `INSERT INTO products (name, code, unit_price) VALUES (?, ?, ?)`
    res, err := repo.DB.Exec(query, p.Name, p.Code, p.UnitPrice)
    if err != nil {
        return err
    }
    
    // Captura o ID gerado e popula a struct
    id, _ := res.LastInsertId()
    p.ID = int(id)
    return nil
}

func (repo *ProductRepository) Update(p models.ProductModel) error {
	query := `UPDATE products SET name = ?, unit_price = ? WHERE code = ?`
	_, err := repo.DB.Exec(query, p.Name, p.UnitPrice, p.Code)
	return err
}

func (repo *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := repo.DB.Exec(query, id)
	return err
}

func (repo *ProductRepository) GetAll() ([]models.ProductModel, error) {
    query := `SELECT id, name, code, unit_price FROM products`
    rows, err := repo.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []models.ProductModel
    for rows.Next() {
        var p models.ProductModel
        if err := rows.Scan(&p.ID, &p.Name, &p.Code, &p.UnitPrice); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}

func (repo *ProductRepository) GetByID(id int) (*models.ProductModel, error) {
	query := `SELECT id, name, code, unit_price FROM products WHERE id = ?`
	row := repo.DB.QueryRow(query, id)

	var p models.ProductModel
	err := row.Scan(&p.ID, &p.Name, &p.Code, &p.UnitPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err 
		}
		return nil, err
	}
	return &p, nil
}

func (repo *ProductRepository) Exists(code string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE code = ?)`
	err := repo.DB.QueryRow(query, code).Scan(&exists)
	return exists, err
}