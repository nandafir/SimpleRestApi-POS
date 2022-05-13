package master

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

func (s Service) SubmitItem(params ItemRequest) error {
	// p := ProductModel{
	// 	ProductName: params.ProductName,
	// 	Category:    params.Category,
	// 	Barcode:     params.Barcode,
	// 	Price:       params.Price,
	// 	Stock:       params.Stock,
	// 	CreatedBy:   "jhon doe",
	// }

	err := s.repository.SubmitItem(params)
	if err != nil {
		utils.ErrorLog("SQL Error on InsertItem", err)
		return err
	}

	return nil
}

func (s Service) GetItems() (ItemModels, error) {
	res, err := s.repository.GetItems()
	if err != nil {
		utils.ErrorLog("SQL Error on GetItems", err)
		return nil, err
	}

	return *res, nil
}
