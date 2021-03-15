/**
 * @Time : 2021/3/15 3:20 下午
 * @Author : MassAdobe
 * @Description: db
**/
package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 1:49 下午
 * @Description: 数据库实体类
**/
var (
	DB *gorm.DB
)

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 3:31 下午
 * @Description: 初始化数据库
**/
func InitDB() {
	if gg, err := gorm.Open("mysql", "root:abcdefg@tcp(127.0.0.1:3306)/home_features?charset=utf8&parseTime=True&loc=Local"); err != nil {
		fmt.Println("【数据库连接】", "连接失败", err)
		os.Exit(1)
	} else {
		gg.DB().SetMaxIdleConns(2)
		gg.DB().SetMaxOpenConns(10)
		gg.Debug()
		gg.LogMode(true)
		if err := gg.DB().Ping(); err != nil {
			fmt.Println("【数据库连接】", "初始化失败", err)
			os.Exit(1)
		} else {
			fmt.Println("【数据库连接】", "初始化成功")
			DB = gg
		}
	}
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/28 2:29 下午
 * @Description: 关停数据库连接池，释放句柄
**/
func CloseDb() {
	if err := DB.Close(); err != nil {
		fmt.Println("【数据库连接】", "关闭数据库读库连接池失败")
	}
}
