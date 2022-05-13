package product

import (
	"time"

	"github.com/google/uuid"
)

type ProductModel struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ProductName string     `db:"product_name" json:"product_name"`
	Category    string     `db:"category" json:"category"`
	Barcode     string     `db:"barcode" json:"barcode"`
	Price       float64    `db:"price" json:"price"`
	Stock       float64    `db:"stock" json:"stock"`
	CreateDate  time.Time  `db:"create_date" json:"create_date"`
	CreatedBy   string     `db:"created_by" json:"created_by"`
	UpdateDate  *time.Time `db:"update_date" json:"update_date"`
	UpdatedBy   *string    `db:"updated_by" json:"updated_by"`
}

type ProductModels []ProductModel
