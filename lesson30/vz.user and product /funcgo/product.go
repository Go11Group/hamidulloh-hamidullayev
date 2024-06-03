package funcgo

import (
	"database/sql"
	"Modul/modul"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductsRepo(DB *sql.DB) *ProductRepo {
	return &ProductRepo{DB: DB}
}

func (c *ProductRepo) ProductsCreate(product modul.Products) error {
	_, err := c.DB.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)",
		product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil {
		return err
	}
	return nil
}

func (c *ProductRepo) ProductsDelete(id uint) error {
	_, err := c.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *ProductRepo) ProductsUpdate(id uint, product modul.Products) error {
	_, err := c.DB.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5",
		product.Name, product.Description, product.Price, product.Stock_quantity, id)
	if err != nil {
		return err
	}
	return nil
}
func (c *ProductRepo) ProductsID(id uint) (*modul.Products, error) {
	produc := &modul.Products{}
	err := c.DB.QueryRow("SELECT id, name, description, price, stock_quantity  FROM products WHERE id = $1", id).
		Scan(&produc.Id, &produc.Name, &produc.Description, &produc.Price, &produc.Stock_quantity)
	if err != nil {
		return nil, err
	}
	return produc, nil
}


func (c *ProductRepo) ProductsSelect() ([]modul.Products, error) {
	rows, err := c.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []modul.Products{}
	for rows.Next() {
		product := modul.Products{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
