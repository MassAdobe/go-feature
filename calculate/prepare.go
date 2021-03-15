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

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 4:19 下午
 * @Description: 准备数据
**/
func PrepareData() {
	whole := mapper.GetWhole()
	for _, v := range whole {
		fmt.Println(v)
	}
}
