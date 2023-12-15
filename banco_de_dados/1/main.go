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

	product2 := NewProduct("Monitor", 1020.90)
	err = insertProduct(db, product2)
	if err != nil{
		panic(err)
	}

	// p, err := selectProduct(db, product.ID)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, possui o preço de R$%.2f\n", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil{
		panic(err)
	}

	for _, p := range products{
		fmt.Printf("Product: %v, possui o preço de R$%.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, product2.ID)
	if err != nil{
		panic(err)
	}
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

func selectAllProducts(db *sql.DB)([]Product, error){
	rows, err := db.Query("select id, name, price from products")
	if err != nil{
		return nil, err
	}

	defer rows.Close()
	var products []Product

	for rows.Next(){
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil{
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}


func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = $1")
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil{
		return err
	}

	return nil
}