package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/williamtrevisan/go-arch-hexagonal/app"
)

type ProductDb struct {
    db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
    return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (app.ProductInterface, error) {
    var product app.Product

    stmt, err := p.db.Prepare("SELECT * FROM products WHERE id = ?")
    if err != nil {
        return nil, err
    }

    err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
    if err != nil {
        return nil, err
    }

    return &product, nil
}

func (p *ProductDb) Save(product app.ProductInterface) (app.ProductInterface, error) {
    var rows int

    p.db.QueryRow("Select id from products where id=?", product.GetID()).Scan(&rows)
    if rows == 0 {
        _, err := p.create(product)
        if err != nil {
            return nil, err
        }
    } else {
        _, err := p.update(product)
        if err != nil {
            return nil, err
        }
    }

    return product, nil
}

func (p *ProductDb) create(product app.ProductInterface) (app.ProductInterface, error) {
    stmt, err := p.db.Prepare(`
        INSERT INTO products(id, name, price, status) 
        VALUES(?, ?, ?, ?)
    `)
    if err != nil {
        return nil, err
    }

    _, err = stmt.Exec(
        product.GetID(),
        product.GetName(),
        product.GetPrice(),
        product.GetStatus(),
    )
    if err != nil {
        return nil, err
    }

    err = stmt.Close()
    if err != nil {
        return nil, err
    }

    return product, nil
}

func (p *ProductDb) update(product app.ProductInterface) (app.ProductInterface, error) {
    _, err := p.db.Exec(
        "update products set name = ?, price=?, status=? where id = ?",
        product.GetName(),
        product.GetPrice(),
        product.GetStatus(),
        product.GetID(),
    )
    if err != nil {
        return nil, err
    }
    
    return product, nil
}
