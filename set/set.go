package set

import (
	"fmt"
	"strings"
)

type any = interface{}

// Set struct
type Set struct {
	data map[any]bool
}

// NewSet return a pointer to a set and init it using the slice
func NewSet(data ...any) *Set {
	s := &Set{
		data: make(map[any]bool),
	}
	for _, d := range data {
		s.Add(d)
	}
	return s
}

// Add the data to the set, return whether add successfully
func (s *Set) Add(data any) bool {
	_, exist := s.data[data]
	if exist {
		return false
	}
	s.data[data] = true
	return true
}

// Exist return the data exist in the set
func (s *Set) Exist(data any) bool {
	_, exist := s.data[data]
	return exist
}

// Clear the data from the set
func (s *Set) Clear(data any) {
	delete(s.data, data)
}

// Copy the set
func (s *Set) Copy() *Set {
	ss := NewSet()
	for data := range s.data {
		ss.Add(data)
	}
	return ss
}

// Len the set
func (s *Set) Len() int {
	return len(s.data)
}

// String of the set
func (s *Set) String() string {
	str := make([]string, s.Len())
	var i int
	for data := range s.data {
		str[i] = fmt.Sprint(data)
		i++
	}
	return fmt.Sprintf("{%s}", strings.Join(str, ", "))
}

// Equal compare two set is equal
func (s *Set) Equal(s2 *Set) bool {
	return Equal(s, s2)
}

// Union two sets: a ∪ b
func (s *Set) Union(s2 *Set) *Set {
	return Union(s, s2)
}

// Intersect two sets: a ∩ b
func (s *Set) Intersect(s2 *Set) *Set {
	return Intersect(s, s2)
}

// Difference two sets: a - b
func (s *Set) Difference(s2 *Set) *Set {
	return Difference(s, s2)
}

// Equal compare two set is equal
func Equal(a *Set, b *Set) bool {
	if a.Len() != b.Len() {
		return false
	}
	for data := range a.data {
		if !b.Exist(data) {
			return false
		}
	}
	return true
}

// Union two sets: a ∪ b
func Union(a *Set, b *Set) *Set {
	s := NewSet()
	for data := range a.data {
		s.Add(data)
	}
	for data := range b.data {
		s.Add(data)
	}
	return s
}

// Intersect two sets: a ∩ b
func Intersect(a *Set, b *Set) *Set {
	s := NewSet()
	for data := range a.data {
		if b.Exist(data) {
			s.Add(data)
		}
	}
	return s
}

// Difference two sets: a - b
func Difference(a *Set, b *Set) *Set {
	s := a.Copy()
	for data := range b.data {
		s.Clear(data)
	}
	return s
}
