package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // postresql driver
)

type Product struct{
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *Product{
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main(){
	db, err := sql.Open("postgres", "postgres://postgres:admin@localhost:5432/app?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 2570.90)
	err = insertProduct(db, product)
	if err != nil{
		panic(err)
	}

	product.Name = "Notebook Gamer"
	err = updateProduct(db, product)
	if err != nil{
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de R$%.2f\n", p.Name, p.Price)

}

func insertProduct(db *sql.DB, product * Product) error{
	stmt, err := db.Prepare("insert into products(id, name, price) values($1, $2, $3)")
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil{
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product * Product) error{
	stmt, err := db.Prepare("update products set name = $1, price = $2 where id = $3")
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil{
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error){
	stmt, err := db.Prepare("select id, name, price from products where id = $1")
	if err != nil{
		return nil, err
	}
	defer stmt.Close()
	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil{
		return nil, err
	}

	return &p, nil

}