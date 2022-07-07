package master

import (
	"anaconda/utils"
	"fmt"

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

	softDeleteItem = `
	UPDATE item set deleted =1
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

func (r repository) SoftDeleteItemByID(id int) error {

	tx, err := r.db.Beginx()
	if err != nil {
		utils.ErrorLog("SoftDelete DeleteItem", err)
		return err
	}
	condition := fmt.Sprintf(" WHERE item_id = %v ", id)
	_, err = tx.Exec(softDeleteItem + condition)
	if err != nil {
		utils.ErrorLog("Repo SoftDeleteItemByID Exec", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r repository) GetItemByID(id int) (*ItemModel, error) {
	dest := &ItemModel{}
	condition := fmt.Sprintf(" WHERE item_id = %v ", id)
	err := r.db.Get(dest, selectItem+condition)
	if err != nil {
		utils.ErrorLog("ProductRepository GetItemByID Select", err)
		return nil, err
	}
	return dest, nil
}
