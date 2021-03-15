/**
 * @Time : 2021/3/15 3:19 下午
 * @Author : MassAdobe
 * @Description: entity
**/
package entity

import "time"

// TGoodPrice 商品价格表
type TGoodPrice struct {
	PriceID     uint32    `gorm:"primary_key;column:price_id;type:int(10) unsigned;not null"` // 商品价格ID
	GoodID      uint32    `gorm:"column:good_id;type:int(10) unsigned;not null"`              // 商品ID
	PriceAmount float64   `gorm:"column:price_amount;type:decimal(10,2);not null"`            // 商品价格(单价)
	PriceType   string    `gorm:"column:price_type;type:char(1);not null"`                    // 价格种类(0:买入;1:卖出;2:保价;3:运费)
	CreatedTm   time.Time `gorm:"column:created_tm;type:timestamp;not null"`                  // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *TGoodPrice) TableName() string {
	return "t_good_price"
}
