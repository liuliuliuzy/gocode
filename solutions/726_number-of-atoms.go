package solutions

import (
	"sort"
	"strconv"
	"unicode"
)

//相关标签里有个 栈
//错误尝试
// func countOfAtoms(formula string) string {
// 	res := ""
// 	//记录原子个数
// 	ht := map[string]int{}
// 	stack := []byte{}
// 	atomIndexs := []string{}
// 	start := 0
// 	for start < len(formula) {
// 		//如果是数字
// 		if formula[start] >= 48 && formula[start] <= 57 {
// 			count := formula[start] - 48
// 			start++
// 			for start < len(formula) && (formula[start] >= 48 && formula[start] <= 57) {
// 				count = count*10 + (formula[start] - 48)
// 				start++
// 			}
// 			atoms := []string{}
// 			//接下来出栈，获取原子
// 			//不可能出现的情况
// 			if len(stack) == 0 {
// 				fmt.Println("Program error.")
// 				return res
// 			}
// 			tmp := ""
// 			for _, c := range stack {
// 				//跳过第一个字符
// 				if c == '(' || c == ')' {
// 					continue
// 				}
// 				if c >= 65 && c <= 90 {
// 					if tmp != "" {
// 						atoms = append(atoms, tmp)
// 					}
// 					tmp = string([]byte{c})
// 				} else {
// 					tmp += string([]byte{c})
// 				}
// 			}
// 			atoms = append(atoms, tmp)
// 			if stack[0] == '(' {
// 				for _, atom := range atoms {
// 					ht[atom] += int(count)
// 				}
// 			} else {
// 				for i, atom := range atoms {
// 					if i == len(atoms)-1 {
// 						ht[atom] += int(count)
// 					} else {
// 						ht[atom] += 1
// 					}
// 				}
// 			}
// 			fmt.Println(atoms, stack)
// 			stack = []byte{}
// 			continue
// 		} else {
// 			stack = append(stack, formula[start])
// 			start++
// 		}
// 	}
// 	//获得了哈希表之后拼凑得到结果
// 	for k, _ := range ht {
// 		atomIndexs = append(atomIndexs, k)
// 	}
// 	sort.Strings(atomIndexs)
// 	for _, index := range atomIndexs {
// 		if ht[index] > 1 {
// 			res += (index + strconv.Itoa(ht[index]))
// 		} else {
// 			res += index
// 		}
// 	}
// 	return res
// }

//坏了：括号里面还可能有数字和括号的嵌套
// func countOfAtoms(formula string) string {
// 	res := ""
// 	//记录原子个数
// 	ht := map[string]int{}
// 	stack := []byte{}
// 	atomIndexs := []string{}
// 	start := 0
// 	for start < len(formula) {
// 		//如果是数字
// 		if formula[start] >= 48 && formula[start] <= 57 {
// 			count := formula[start] - 48
// 			start++
// 			for start < len(formula) && (formula[start] >= 48 && formula[start] <= 57) {
// 				count = count*10 + (formula[start] - 48)
// 				start++
// 			}
// 			//处理栈
// 			i := len(stack) - 1
// 			tmp := ""
// 			//如果是'(....)'的情况
// 			if stack[i] == ')' {
// 				i--
// 				for i > -1 && stack[i] != '(' {
// 					//从后往前遍历，碰到大写字母
// 					if stack[i] >= 65 && stack[i] <= 90 {
// 						tmp = string([]byte{stack[i]}) + tmp
// 						ht[tmp] += int(count)
// 						tmp = ""
// 					} else {
// 						tmp = string([]byte{stack[i]}) + tmp
// 					}
// 					i--
// 				}
// 				// ht[tmp] += int(count)
// 			} else {
// 				for i > -1 && (stack[i] >= 97 && stack[i] <= 122) {
// 					tmp = string([]byte{stack[i]}) + tmp
// 					i--
// 				}
// 				tmp = string([]byte{stack[i]}) + tmp
// 				ht[tmp] += int(count)

// 			}
// 			stack = stack[:i]
// 			continue
// 		} else {
// 			stack = append(stack, formula[start])
// 			start++
// 		}
// 	}
// 	tmp := ""
// 	for _, c := range stack {
// 		if c == '(' {
// 			continue
// 		}
// 		if c >= 65 && c <= 90 {
// 			if tmp != "" {
// 				ht[tmp] += 1
// 			}
// 			tmp = string([]byte{c})
// 		} else {
// 			tmp += string([]byte{c})
// 		}
// 	}
// 	ht[tmp] += 1

// 	fmt.Println(ht)
// 	//获得了哈希表之后拼凑得到结果
// 	for k := range ht {
// 		atomIndexs = append(atomIndexs, k)
// 	}
// 	sort.Strings(atomIndexs)
// 	for _, index := range atomIndexs {
// 		if ht[index] > 1 {
// 			res += (index + strconv.Itoa(ht[index]))
// 		} else {
// 			res += index
// 		}
// 	}
// 	return res
// }

//直接搬运官方题解
func countOfAtoms(formula string) string {
	i, n := 0, len(formula)

	parseAtom := func() string {
		start := i
		i++ // 扫描，跳过首字母
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++ // 扫描首字母后的小写字母
		}
		return formula[start:i]
	}

	parseNum := func() (num int) {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1 // 不是数字，视作 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = num*10 + int(formula[i]-'0') // 扫描数字
		}
		return
	}

	stk := []map[string]int{{}}
	for i < n {
		if ch := formula[i]; ch == '(' {
			i++
			stk = append(stk, map[string]int{}) // 将一个空的哈希表压入栈中，准备统计括号内的原子数量
		} else if ch == ')' {
			i++
			num := parseNum() // 括号右侧数字
			atomNum := stk[len(stk)-1]
			stk = stk[:len(stk)-1] // 弹出括号内的原子数量
			for atom, v := range atomNum {
				stk[len(stk)-1][atom] += v * num // 将括号内的原子数量乘上 num，加到上一层的原子数量中
			}
		} else {
			atom := parseAtom()
			num := parseNum()
			stk[len(stk)-1][atom] += num // 统计原子数量
		}
	}

	atomNum := stk[0]
	type pair struct {
		atom string
		num  int
	}
	pairs := make([]pair, 0, len(atomNum))
	for k, v := range atomNum {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })

	ans := []byte{}
	for _, p := range pairs {
		ans = append(ans, p.atom...)
		if p.num > 1 {
			ans = append(ans, strconv.Itoa(p.num)...)
		}
	}
	return string(ans)
}

func CountOfAtoms(formula string) string {
	return countOfAtoms(formula)
}
