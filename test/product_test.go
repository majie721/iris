package test

import (
	"iris_stu/common"
	"iris_stu/datamodels"
	"iris_stu/repositories"
	"testing"
)

func Test_insert(t *testing.T)  {
	mysql,err := common.NewMysqlConn()
	if err !=nil{
		t.Fatal(err)
	}
	rep :=  repositories.NewProductManager("iris_product",mysql)
	insertModel := &datamodels.Product{
		Name: "华为",
		Num: 100,
		ProductImage: "ProductImage",
		ProductUrl: "ProductUrl",
	}

	id,err := rep.Insert(insertModel)
	if err != nil{
		t.Fatal(err)
	}else{
		t.Logf("插入成功,ID: %d,data:%+v",id,insertModel)
	}
}


