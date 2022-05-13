package product

import (
	"anaconda/utils"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

var (
	insertProduct = `
	INSERT INTO
		products (
			product_name,
			category,
			barcode,
			price,
			stock,
			create_date,
			created_by
		)
	VALUES
		(
			:product_name,
			:category,
			:barcode,
			:price,
			:stock,
			NOW(),
			:created_by
		)
	`

	selectProduct = `
	SELECT
		id,
		product_name,
		category,
		barcode,
		price,
		stock,
		create_date,
		update_date,
		updated_by
	FROM
		products p
	`
)

func (r repository) SubmitProduct(param ProductModel) error {
	tx, err := r.db.Beginx()
	if err != nil {
		utils.ErrorLog("SubmitSurvey InsertProduct", err)
		return err
	}

	_, err = tx.NamedExec(insertProduct, param)
	if err != nil {
		utils.ErrorLog("SurveyRepository UpsertSurvey NamedExec Survey", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r repository) GetProducts() (*ProductModels, error) {
	dest := &ProductModels{}
	err := r.db.Select(dest, selectProduct+" ORDER BY product_name ASC")
	if err != nil {
		utils.ErrorLog("ProductRepository GetProdcut Select", err)
		return nil, err
	}
	return dest, nil
}
