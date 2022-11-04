package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.90)

	// CREATE
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// READ
	// SELECT ALL
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Println(p)
	}

	// SELECT ONE
	p, err := selectOneProduct(db, products[0].ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)

	product.Price = 2000.0

	// UPDATE
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// DELETE
	deleteProduct(db, p.ID)

}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name=?, price=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func deleteProduct(db *sql.DB, productId string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(productId)
	if err != nil {
		return err
	}

	return nil
}

func selectOneProduct(db *sql.DB, productId string) (*Product, error) {
	stmt, err := db.Prepare("SELECT * FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product

	err = stmt.QueryRow(productId).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, err
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
