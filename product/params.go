package product

type ProductRequest struct {
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
}

type ProductsRequest []ProductRequest
