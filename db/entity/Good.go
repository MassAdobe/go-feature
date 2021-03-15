/**
 * @Time : 2021/3/15 3:14 下午
 * @Author : MassAdobe
 * @Description: entity
**/
package entity

import "time"

// TGood 商品表
type TGood struct {
	GoodID    uint32    `gorm:"primary_key;column:good_id;type:int(10) unsigned;not null"` // 商品ID
	GoodName  string    `gorm:"column:good_name;type:varchar(32);not null"`                // 商品名称
	CreatedTm time.Time `gorm:"column:created_tm;type:timestamp;not null"`                 // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *TGood) TableName() string {
	return "t_good"
}
