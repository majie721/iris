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
		Name: "三星2",
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

func Test_update(t *testing.T)  {
	mysql,err := common.NewMysqlConn()
	if err !=nil{
		t.Fatal(err)
	}
	rep :=  repositories.NewProductManager("iris_product",mysql)
	insertModel := &datamodels.Product{
		ID: 5,
		Name: "三星A",
		Num: 100,
		ProductImage: "ProductImage",
		ProductUrl: "ProductUrl",
	}

	err = rep.Update(insertModel)
	if err != nil{
		t.Fatal(err)
	}else{
		t.Logf("update 成功")
	}
}


func Test_delete(t *testing.T)  {
	mysql,err := common.NewMysqlConn()
	if err !=nil{
		t.Fatal(err)
	}
	rep :=  repositories.NewProductManager("iris_product",mysql)
	res := rep.Delete(6)
	if res == false{
		t.Fatal(err)
	}else{
		t.Logf("delete 成功")
	}
}



func Test_select_row(t *testing.T)  {
	mysql,err := common.NewMysqlConn()
	if err !=nil{
		t.Fatal(err)
	}
	rep :=  repositories.NewProductManager("iris_product",mysql)
	modle,err := rep.SelectByKey(1)
	if err != nil{
		t.Fatal(err)
	}else{
		if modle == nil{
			t.Logf("select 为空")
		}else{
			t.Logf("%+v",modle)
		}
	}
}


