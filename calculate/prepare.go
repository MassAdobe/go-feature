/**
 * @Time : 2021/3/15 4:19 下午
 * @Author : MassAdobe
 * @Description: calculate
**/
package calculate

import (
	"fmt"
	"github.com/MassAdobe/go-feature/db/mapper"
)

type Calculate struct {
	GoodName          string  // 商品名称
	PriceAmount       float64 // 商品单价
	PriceType         string  // 价格种类(0:买入;1:卖出;2:保价;3:运费)
	TransactionType   string  // 交易种类(0:买入;1:卖出;2:补仓)
	TransactionDate   string  // 交易日期
	TransactionAmount int     // 交易数量
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 4:19 下午
 * @Description: 准备数据
**/
func PrepareData() {
	whole, calMap := mapper.GetWhole(), make(map[uint64]map[string]*Calculate)
	for _, entity := range whole {
		if _, okay := calMap[entity.GoodId]; okay {
			if "1" == entity.TransactionType { // 卖出
				calMap[entity.GoodId]["+"] = &Calculate{
					entity.GoodName,
					entity.PriceAmount,
					priceType(entity.PriceType),
					transactionType(entity.TransactionType),
					entity.TransactionDate,
					entity.TransactionAmount,
				}
			} else {
				calMap[entity.GoodId]["-"] = &Calculate{
					entity.GoodName,
					entity.PriceAmount,
					priceType(entity.PriceType),
					transactionType(entity.TransactionType),
					entity.TransactionDate,
					entity.TransactionAmount,
				}
			}
		} else {
			calMap[entity.GoodId] = make(map[string]*Calculate)
			if "1" == entity.TransactionType { // 卖出
				calMap[entity.GoodId]["+"] = &Calculate{
					entity.GoodName,
					entity.PriceAmount,
					priceType(entity.PriceType),
					transactionType(entity.TransactionType),
					entity.TransactionDate,
					entity.TransactionAmount,
				}
			} else {
				calMap[entity.GoodId]["-"] = &Calculate{
					entity.GoodName,
					entity.PriceAmount,
					priceType(entity.PriceType),
					transactionType(entity.TransactionType),
					entity.TransactionDate,
					entity.TransactionAmount,
				}
			}
		}
	}

	fmt.Println()
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 4:43 下午
 * @Description: 价格种类(0:买入;1:卖出;2:保价;3:运费)
**/
func priceType(typ string) string {
	if "0" == typ {
		return "买入"
	} else if "1" == typ {
		return "卖出"
	} else if "2" == typ {
		return "保价"
	} else {
		return "运费"
	}
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 4:45 下午
 * @Description: 交易种类(0:买入;1:卖出;2:补仓)
**/
func transactionType(typ string) string {
	if "0" == typ {
		return "买入"
	} else if "1" == typ {
		return "卖出"
	} else {
		return "补仓"
	}
}
