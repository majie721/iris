package repositories

import (
	"database/sql"
	"iris_stu/common"
	"iris_stu/datamodels"
	"strconv"
)

type IProduct interface {
	Conn() error
	Insert(product *datamodels.Product)(int64,error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64)(*datamodels.Product,error)
	SelectAll()([]*datamodels.Product,error)
}

type ProductManager struct {  //实现接口
	table string
	mysqlConn *sql.DB
}


func NewProductManager(table string,db *sql.DB) IProduct {
	return &ProductManager{table:table,mysqlConn: db}
}

func (p *ProductManager ) Conn()(err error)  {
	if p.mysqlConn == nil{
		mysql,err := common.NewMysqlConn()
		if err !=nil{
			return  err
		}
		p.mysqlConn = mysql
	}
	return nil
}

func (p *ProductManager) Insert(product *datamodels.Product)(id int64, err error)  {
	if err := p.Conn();err !=nil{
		return 0, err
	}

	sql := "insert iris_product set name=?,num=?,product_image=?,product_url=?"
	stmt,err :=p.mysqlConn.Prepare(sql)
	if err!=nil{
		return 0,err
	}

	result,err := stmt.Exec(product.Name,product.Num,product.ProductImage,product.ProductUrl)
	if err !=nil{
		return 0,err
	}

	return result.LastInsertId()

}

func (p ProductManager) Update(product *datamodels.Product) (err error) {
	if err := p.Conn();err !=nil{
		return err
	}

	sql := "update product set name=?,num=?,product_image=?,product_url=? where id="+strconv.FormatInt(product.ID,10)
	stmt,err :=p.mysqlConn.Prepare(sql)
	if err!=nil{
		return err
	}

	_,err = stmt.Exec(product.Name,product.Num,product.ProductImage,product.ProductUrl)
	return err
}


func (p *ProductManager) Delete(id int64 ) bool {
	if err := p.Conn();err !=nil{
		return false
	}

	sql := "delete from product where id=?"

	stmt,err := p.mysqlConn.Prepare(sql)
	if err !=nil{
		return  false
	}

	_,err = stmt.Exec(id)
	if err !=nil{
		return  false
	}

	return true
}

func (p ProductManager) SelectByKey(int64)(res *datamodels.Product,err error) {
	if err := p.Conn();err !=nil{
		return &datamodels.Product{},err
	}

	sql := "select * from product where id=?"

	row,err := p.mysqlConn.Query(sql)
	if err !=nil{
		return &datamodels.Product{},err
	}
	defer row.Close()
	result := common.GetResultRow(row)
	if len(result) == 0{
		return &datamodels.Product{},nil
	}

	common.DataToStructByTagSql(result,res)

	return
}

func (p ProductManager) SelectAll()(productArr []*datamodels.Product,err error) {
	if err := p.Conn();err !=nil{
		return nil,err
	}


	sql := "select * from product "

	rows,err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err !=nil{
		return nil,err
	}


	result := common.GetResultRows(rows)
	if len(result) == 0{
		return nil,nil
	}

	for _,v := range result{
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v,product)
		productArr = append(productArr, product)
	}
	return
}


