package models

import "fmt"

type Product struct {
	Id uint64
	Name string
	Price float64
	Quantity int
	Amount float64
	Category Category
}

func GetProducts() ([]Product, error) {
	con := Connect()
	defer con.Close()
	sql := `SELECT c.id, c.description, 
			p.id, p.name, p.price, p.quantity, p.amount
			FROM products AS p 
			INNER JOIN category AS c 
			ON c.id = p.category`
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var products []Product
	for rs.Next() {
		var product Product
		err := rs.Scan(&product.Category.Id,
					   &product.Category.Description,
					   &product.Id,
					   &product.Name,
					   &product.Price,
					   &product.Quantity,
					   &product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func SearchProducts(search string) ([]Product, error) {
	search = fmt.Sprintf("%%%s%%", search)
	con := Connect()
	defer con.Close()
	sql := `SELECT c.id, c.description, 
			p.id, p.name, p.price, p.quantity, p.amount
			FROM products AS p 
			INNER JOIN category AS c 
			ON c.id = p.category
			WHERE c.description LIKE $1 OR p.name LIKE $2`
	stmt, err := con.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rs, err := stmt.Query(search, search)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var products []Product
	for rs.Next() {
		var product Product
		err := rs.Scan(&product.Category.Id,
					   &product.Category.Description,
					   &product.Id,
					   &product.Name,
					   &product.Price,
					   &product.Quantity,
					   &product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}