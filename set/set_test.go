package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s1 := NewSet()
	if l := s1.Len(); l != 0 {
		t.Fatalf("Empty set has length %d\n", l)
	}
	if got := s1.Add("a"); got != true {
		t.Fatalf("Add() returns %v\n", got)
	}
	if got := s1.Add(2); got != true {
		t.Fatalf("Add() returns %v\n", got)
	}
	if got := s1.Add('c'); got != true {
		t.Fatalf("Add() returns %v\n", got)
	}
	if temp := NewSet("a", 2, 'c'); !s1.Equal(temp) {
		t.Fatalf("Set %s is not equal with %v\n", s1.String(), temp)
	}
	if l := s1.Len(); l != 3 {
		t.Fatalf("Set %s length is %d\n", s1.String(), l)
	}

	s2 := s1.Copy()
	if got := s2.Add('c'); got != false {
		t.Fatalf("Add() returns %v\n", got)
	}

	if !s1.Exist("a") {
		t.Fatalf("Set %s doesn't has a\n", s1.String())
	}
	s1.Clear("a")
	if s1.Exist("a") {
		t.Fatalf("Set %s does has a\n", s1.String())
	}

	if temp := s1.Union(s2); !temp.Equal(s2) {
		t.Fatalf("Union is %s\n", temp.String())
	}
	if temp := s1.Intersect(s2); !temp.Equal(s1) {
		t.Fatalf("Intersection is %s\n", temp.String())
	}
	if temp := s2.Difference(s1); temp.String() != "{a}" {
		t.Fatalf("Difference is %s\n", temp.String())
	}

	if a, b := NewSet(), NewSet("a"); a.Equal(b) {
		t.Fatalf("%s and %s is equal\n", a.String(), b.String())
	}

	if a, b := NewSet("b"), NewSet("a"); a.Equal(b) {
		t.Fatalf("%s and %s is equal\n", a.String(), b.String())
	}
}
