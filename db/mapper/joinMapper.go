/**
 * @Time : 2021/3/15 3:55 下午
 * @Author : MassAdobe
 * @Description: mapper
**/
package mapper

import (
	"fmt"
	"github.com/MassAdobe/go-feature/db"
	"github.com/MassAdobe/go-feature/db/entity"
	"os"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:26 下午
 * @Description: 获取所有项目
**/
func GetWhole() []*entity.WholeEntity {
	sql := `
select c.good_id            as good_id,
       c.good_name          as good_name,
       b.price_amount       as price,
       b.price_type         as price_type,
       a.transaction_type   as transaction_type,
       a.transaction_date   as date,
       a.transaction_amount as amount
from t_good_transaction a
         left join t_good_price b on a.price_id = b.price_id
         left join t_good c on b.good_id = c.good_id
order by c.good_id;
`
	rtn := make([]*entity.WholeEntity, 0)
	if err := db.DB.Raw(sql).Scan(&rtn).Error; err != nil {
		fmt.Println("【数据库查询错误】", "获取所有项目")
		os.Exit(1)
	}
	return rtn
}
