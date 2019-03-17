package ints

type Set struct {
	m map[int]int
}

func NewSetFromSlice(nums []int) *Set {
	s := new(Set)
	for _, n := range nums {
		s.m[n]++
	}
	return s
}

func (s *Set) Size() int {
	return len(s.m)
}
