package cache

import (
	"fmt"
	"testing"
)

func TestSeckillGoodStockKey(t *testing.T) {
	key := SeckillGoodStockKey(99)
	fmt.Println(key)
}
