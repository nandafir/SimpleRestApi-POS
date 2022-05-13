package master

type ItemModel struct {
	ItemID     int     `db:"item_id" json:"item_id"`
	ItemName   string  `db:"item_name" json:"item_name"`
	ItemSize   string  `db:"item_size" json:"item_size"`
	CategoryID int     `db:"category_id" json:"category_id"`
	SuplierID  int     `db:"suplier_id" json:"suplier_id"`
	Barcode    string  `db:"barcode" json:"barcode"`
	Price      float64 `db:"price" json:"price"`
	Stock      float64 `db:"stock" json:"stock"`
}

type ItemModels []ItemModel
