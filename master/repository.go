package master

import (
	"anaconda/utils"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

var (
	insertItem = `
	insert into item
		(item_name,
		item_size ,
		category_id ,
		suplier_id ,
		barcode ,
		price)
	values
		(:item_name,
		:item_size ,
		:category_id ,
		:suplier_id ,
		:barcode ,
		:price)
	`

	selectItem = `
	SELECT		
		item_id,
		item_name,
		item_size,
		category_id,
		suplier_id,
		barcode,
		price,
		stock
	FROM
		item
	`
)

func (r repository) SubmitItem(param ItemRequest) error {
	tx, err := r.db.Beginx()
	if err != nil {
		utils.ErrorLog("SubmitItem InsertItem", err)
		return err
	}

	_, err = tx.NamedExec(insertItem, param)
	if err != nil {
		utils.ErrorLog("MasterRepository NamedExec item", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r repository) GetItems() (*ItemModels, error) {
	dest := &ItemModels{}
	err := r.db.Select(dest, selectItem+" ORDER BY item_name ASC")
	if err != nil {
		utils.ErrorLog("ProductRepository GetItem Select", err)
		return nil, err
	}
	return dest, nil
}
