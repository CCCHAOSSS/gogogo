package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
* @Author: ccchaosss@github.com
* @Date: 2024/9/16 21:12
* @Des: 数据库操作
 */

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type APIProduct struct {
	Code  string
	Price uint
}

func main() {
	dsn := "root:12345abc@tcp(127.0.0.1:3306)/temp_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//迁移schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	//var product Product

	//err = db.Create(&Product{Code: "D01", Price: 100}).Error
	//if err != nil {
	//	panic(err)
	//}
	//
	////read
	//var product Product
	//err = db.First(&product, 1).Error ////根据整形主键查找
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(product)
	//}
	//db.First(&product, "code = ?", "D42")
	//fmt.Println(product) //根据code查找

	//查询
	//db.First(&product) //第一条记录（主键升序的第一条）
	//db.Take() //获取一条，没有指定排序
	//db.Last(&product) //最后一条记录（主键降序的第一条）

	//var products []Product
	//db.Where("code = ?", "D42").Find(&products)
	//fmt.Printf("products: %+v\n", products)

	//可以查询指定字段，而不是select *
	var res []APIProduct
	db.Model(&Product{}).Find(&res)
	fmt.Printf("result: %+v\n", res)

}
