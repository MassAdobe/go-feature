/**
 * @Time : 2021/3/15 3:51 下午
 * @Author : MassAdobe
 * @Description: entity
**/
package entity

type WholeEntity struct {
	GoodId            uint64  `gorm:"column:good_id;type:int(11);not null"`
	GoodName          string  `gorm:"column:good_name;type:varchar(32);not null"`
	PriceAmount       float64 `gorm:"column:price;type:decimal(10,2);not null"`
	PriceType         string  `gorm:"column:price_type;type:char(1);not null"`       // 价格种类(0:买入;1:卖出;2:保价;3:运费)
	TransactionType   string  `gorm:"column:transaction_type;type:char(1);not null"` // 交易种类(0:买入;1:卖出;2:补仓)
	TransactionDate   string  `gorm:"column:date;type:date;not null"`                // 交易日期
	TransactionAmount int     `gorm:"column:amount;type:int(11);not null"`           // 交易数量
}
