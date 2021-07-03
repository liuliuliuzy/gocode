package solutions

import "fmt"

func frequencySort(s string) string {
	res := ""
	type pair struct {
		key   byte
		count int
	}
	counts := []pair{}
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] += 1
	}
	for k, v := range mp {
		tmp := pair{k, v}
		counts = append(counts, tmp)
	}
	//开始排序 quick sort
	var qcSort func(int, int)
	qcSort = func(left, right int) {
		if left < right {
			p := counts[left]
			start := left
			end := right
			for start < end {
				for start < end && counts[end].count >= p.count {
					end--
				}
				if start < end {
					counts[start] = counts[end]
					start++
				}
				for start < end && counts[start].count <= p.count {
					start++
				}
				if start < end {
					counts[end] = counts[start]
				}
			}
			//此时start就是分界线，start索引处已经放置了正确的元素
			counts[start] = p
			qcSort(left, start-1)
			qcSort(start+1, right)
		}
	}
	qcSort(0, len(counts)-1)
	fmt.Println(counts)
	for i := len(counts) - 1; i > -1; i-- {
		tmpC := []byte{counts[i].key}
		tmpS := string(tmpC)
		for counts[i].count > 0 {
			res += tmpS
			counts[i].count--
		}
	}
	return res
}

func FrequencySort(s string) string {
	return frequencySort(s)
}
