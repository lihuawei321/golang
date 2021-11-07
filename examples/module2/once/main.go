package main

import (
	"fmt"
	"sync"
)
//单例模式，一块代码希望只初始化一次
type SliceNum []int

func NewSlice() SliceNum {
	return make(SliceNum, 0)

}

func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	return s
}

func main() {
	var once sync.Once
	s := NewSlice()
	// 看源代码理解once的行为
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(16)
	})
}