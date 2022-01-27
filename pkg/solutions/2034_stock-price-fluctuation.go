package solutions

// 使用GoDS的红黑树数据结构
import rbt "github.com/emirpasic/gods/trees/redblacktree"

type StockPrice struct {
	prices       *rbt.Tree   // 有序集合
	timePriceMap map[int]int // 哈希表
	maxTimestamp int         // 最新时间戳
}

func sConstructor() StockPrice {
	return StockPrice{rbt.NewWithIntComparator(), map[int]int{}, 0}
}

/*
难点在于如何在更新新的股价的同时更新最大或最小值
*/
func (sp *StockPrice) Update(timestamp int, price int) {
	if prevPrice := sp.timePriceMap[timestamp]; prevPrice > 0 {
		if times, _ := sp.prices.Get(prevPrice); times.(int) > 1 {
			sp.prices.Put(prevPrice, times.(int)-1)
		} else {
			sp.prices.Remove(prevPrice)
		}
	}
	times := 0
	if val, ok := sp.prices.Get(price); ok {
		times = val.(int)
	}
	sp.prices.Put(price, times+1)
	sp.timePriceMap[timestamp] = price
	if timestamp >= sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	return sp.prices.Right().Key.(int) // 返回值并进行类型断言
}

func (sp *StockPrice) Minimum() int {
	return sp.prices.Left().Key.(int)
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
