/**
 * @Time : 2021/3/15 3:19 下午
 * @Author : MassAdobe
 * @Description: entity
**/
package entity

import "time"

// TGoodTransaction 商品交易表
type TGoodTransaction struct {
	TransactionID     uint32    `gorm:"primary_key;column:transaction_id;type:int(10) unsigned;not null"` // 商品交易ID
	PriceID           uint32    `gorm:"column:price_id;type:int(10) unsigned;not null"`                   // 交易商品的价格
	TransactionAmount int       `gorm:"column:transaction_amount;type:int(11);not null"`                  // 交易数量
	TransactionType   string    `gorm:"column:transaction_type;type:char(1);not null"`                    // 交易种类(0:买入;1:卖出;2:补仓)
	TransactionDate   time.Time `gorm:"column:transaction_date;type:date;not null"`                       // 交易日期
	CreatedTm         time.Time `gorm:"column:created_tm;type:timestamp;not null"`                        // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *TGoodTransaction) TableName() string {
	return "t_good_transaction"
}
