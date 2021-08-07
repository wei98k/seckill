package cache

import "fmt"

func SeckillGoodStockKey(id int64) string {
	return fmt.Sprintf("SECKILL:GOODS:%d:STOCK", id)
}

func SeckillGoodOverKey(id int64) string {
	return fmt.Sprintf("SECKILL:GOODS:%d:OVER", id)
}