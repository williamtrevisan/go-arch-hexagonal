package app

import (
    "errors"
    "github.com/asaskevich/govalidator"
    uuid "github.com/satori/go.uuid"
)

func init() {
    govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
    Disable() error
    Enable() error
    GetID() string
    GetName() string
    GetPrice() float64
    GetStatus() string
    IsValid() (bool, error)
}

const (
    DISABLED = "disabled"
    ENABLED  = "enabled"
)

type Product struct {
    ID     string  `valid:"uuidv4"`
    Name   string  `valid:"required"`
    Price  float64 `valid:"float,optional"`
    Status string  `valid:"required"`
}

func NewProduct() *Product {
    product := &Product{
        ID:     uuid.NewV4().String(),
        Status: DISABLED,
    }

    return product
}

func (p *Product) Disable() error {
    if p.Price == 0 {
        p.Status = DISABLED

        return nil
    }

    return errors.New("The price must be zero in order to have the product disabled.")
}

func (p *Product) Enable() error {
    if p.Price > 0 {
        p.Status = ENABLED

        return nil
    }

    return errors.New("The price must be greater than zero to enable the product.")
}

func (p *Product) GetID() string {
    return p.ID
}

func (p *Product) GetName() string {
    return p.Name
}

func (p *Product) GetPrice() float64 {
    return p.Price
}

func (p *Product) GetStatus() string {
    return p.Status
}

func (p *Product) IsValid() (bool, error) {
    if p.Status == "" {
        p.Status = DISABLED
    }

    if p.Status != ENABLED && p.Status != DISABLED {
        return false, errors.New("The status must be enabled or disabled.")
    }

    if p.Price < 0 {
        return false, errors.New("The price must be greater or equal zero.")
    }

    _, err := govalidator.ValidateStruct(p)
    if err != nil {
        return false, err
    }

    return true, nil
}
