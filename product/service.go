package product

import (
	"anaconda/database"
	"anaconda/utils"
)

type Service struct {
	repository *repository
}

func NewService() *Service {
	return &Service{
		repository: &repository{
			db: database.New(),
		},
	}
}

func (s Service) SubmitProduct(params ProductRequest) error {
	p := ProductModel{
		ProductName: params.ProductName,
		Category:    params.Category,
		Barcode:     params.Barcode,
		Price:       params.Price,
		Stock:       params.Stock,
		CreatedBy:   "jhon doe",
	}

	err := s.repository.SubmitProduct(p)
	if err != nil {
		utils.ErrorLog("SQL Error on InsertProduct", err)
		return err
	}

	return nil
}

func (s Service) GetProducts() (ProductModels, error) {
	res, err := s.repository.GetProducts()
	if err != nil {
		utils.ErrorLog("SQL Error on GetProduct", err)
		return nil, err
	}

	return *res, nil
}
