/**
 * @Time : 2021/3/15 3:39 下午
 * @Author : MassAdobe
 * @Description: go_feature
**/
package main

import (
	"github.com/MassAdobe/go-feature/calculate"
	"github.com/MassAdobe/go-feature/db"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 1:49 下午
 * @Description: 预热项
**/
func init() {
	db.InitDB() // 初始化DB
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 3:42 下午
 * @Description: 主程
**/
func main() {
	calculate.PrepareData()
	db.CloseDb() // 关闭DB
}
