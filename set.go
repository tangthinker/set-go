package set_go

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(element ...T) {
	for _, element := range element {
		s.elements[element] = struct{}{}
	}
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.elements))
	for element := range s.elements {
		elements = append(elements, element)
	}
	return elements
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := NewSet[T]()
	for element := range s.elements {
		unionSet.Add(element)
	}
	for element := range other.elements {
		unionSet.Add(element)
	}
	return unionSet
}

func (s *Set[T]) Unions(others ...*Set[T]) *Set[T] {
	unionSet := NewSet[T]()
	for element := range s.elements {
		unionSet.Add(element)
	}
	for _, other := range others {
		for element := range other.elements {
			unionSet.Add(element)
		}
	}
	return unionSet
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersectionSet := NewSet[T]()
	for element := range s.elements {
		if other.Contains(element) {
			intersectionSet.Add(element)
		}
	}
	return intersectionSet
}

func (s *Set[T]) Intersections(others ...*Set[T]) *Set[T] {
	intersectionSet := NewSet[T]()
	for element := range s.elements {
		intersectionSet.Add(element)
	}
	for _, other := range others {
		intersectionSet = intersectionSet.Intersection(other)
	}
	return intersectionSet
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	differenceSet := NewSet[T]()
	for element := range s.elements {
		if !other.Contains(element) {
			differenceSet.Add(element)
		}
	}
	return differenceSet
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for element := range s.elements {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

func (s *Set[T]) String() string {
	builder := strings.Builder{}

	builder.WriteString("(")

	for i, element := range s.Elements() {
		builder.WriteString(fmt.Sprintf("%v", element))
		if i < len(s.elements)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(")")
	return builder.String()
}
