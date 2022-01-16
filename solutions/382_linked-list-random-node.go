package solutions

import "math/rand"

// import "math/rand"
//
// type Solution struct {
// 	length   int
// 	elements []int
// }
//
// func Constructor(head *ListNode) Solution {
// 	s := Solution{length: 0, elements: []int{}}
// 	for head != nil {
// 		s.elements = append(s.elements, head.Val)
// 		s.length++
// 		head = head.Next
// 	}
// 	return s
// }
//
// func (s *Solution) GetRandom() int {
// 	// generate random step
// 	sample := rand.Intn(s.length)
// 	return s.elements[sample]
// }

type Solution struct {
	lkst *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (s *Solution) GetRandom() int {
	// 水塘抽样
	help := s.lkst
	index := 2
	value := help.Val
	for help.Next != nil {
		help = help.Next
		index++
		r := rand.Intn(index)
		if r == 0 {
			value = help.Val
		}
	}
	return value
}
