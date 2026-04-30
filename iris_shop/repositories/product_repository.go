package repositories

import (
	"database/sql"
	"iris_shop/common"
	"iris_shop/datamodels"
	"strconv"
)

type ProductRepository interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectById(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductRepositoryImpl struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductRepositoryImpl(table string, db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{
		table:     table,
		mysqlConn: db,
	}
}

func (p *ProductRepositoryImpl) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

func (p *ProductRepositoryImpl) Insert(product *datamodels.Product) (productId int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "INSERT " + p.table + " SET product_name=?,product_num=?,product_image=?,product_url=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

func (p *ProductRepositoryImpl) Delete(productId int64) (result bool) {
	if err := p.Conn(); err != nil {
		return
	}
	sql := "DELETE FROM " + p.table + " WHERE id=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(productId)
	if err != nil {
		return
	}
	return true
}

func (p *ProductRepositoryImpl) Update(product *datamodels.Product) (err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "UPDATE " + p.table + " SET product_name=?,product_num=?,product_image=?,product_url=? WHERE id=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	return
}

func (p *ProductRepositoryImpl) SelectById(productId int64) (product *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "SELECT * FROM " + p.table + " WHERE id=" + strconv.FormatInt(productId, 10)
	row, err := p.mysqlConn.Query(sql)
	if err != nil {
		return
	}
	defer row.Close()
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return nil, nil
	}
	common.DataToStructByTagSql(result, product)
	return
}

func (p *ProductRepositoryImpl) SelectAll() (products []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sql := "SELECT * FROM " + p.table
	rows, err := p.mysqlConn.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return
	}
	for _, row := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(row, product)
		products = append(products, product)
	}
	return
}
