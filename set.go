package setlib

import (
	"fmt"
	"reflect"
	"sync"
	"strings"
)

type Set[T comparable] struct {
	mx sync.Mutex
	elementIsStruct bool
	elements map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	s := Set[T]{}
	s.elements = make(map[T]bool)
	s.elementIsStruct = reflect.ValueOf(*new(T)).Kind() == reflect.Struct
	return &s
}

func (s *Set[T]) String() string {
	elemStr := ""
	index := s.size()
	for value := range s.elements {
		valueStr := fmt.Sprintf("%v", value)
		if s.elementIsStruct {
			valueStr = strings.TrimPrefix(valueStr, "{")
			valueStr = strings.TrimSuffix(valueStr, "}")
			valueStr = "(" + valueStr + ")"
		}
		elemStr += valueStr
		if index > 1 {
			elemStr += ", "
		}
		index--
	}
	return fmt.Sprint("{", elemStr, "}")
}

func (s1 *Set[T]) size() int {
	return len(s1.elements)
}

func (s1 *Set[T]) equals(s2 *Set[T]) bool {	
	if s1.size() != s2.size() {
		return false
	}
	if s1.elementIsStruct {
		s2_elements := make([]T, 0)
		for element := range s2.elements {
			s2_elements = append(s2_elements, element)
		}
		done := make([]int, 0)
		for element1 := range s1.elements { 
			for i, element2 := range s2_elements {
				if reflect.DeepEqual(element1, element2) {
					done = append(done, i)
					break
				}
			}
		}
		return len(done) == s2.size()
	} else {
		for element := range s2.elements {
			if !s1.elements[element] {
				return false
			}
		}
	}
	return true
}

func (s1 *Set[T]) Equals(s2 *Set[T]) bool {
	s1.lock(); defer s1.unlock()
	s2.lock(); defer s2.unlock()
	return s1.equals(s2)
}

func (s *Set[T]) lock() {
	s.mx.Lock()
}

func (s *Set[T]) unlock() {
	s.mx.Unlock()
}

func (s *Set[T]) add(element T) {
	s.elements[element] = true
}

func (s *Set[T]) Add(element T) {
	s.lock(); defer s.unlock()
	s.add(element)
}

func (s *Set[T]) remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Remove(element T) {
	s.lock(); defer s.unlock()
	s.remove(element)
}

func (s *Set[T]) Contains(element T) bool {
	s.lock(); defer s.unlock()
	return s.elements[element]
}

func (s *Set[T]) Elements() (elems []T) {
	s.lock(); defer s.unlock()
	
	for element := range s.elements {
		elems = append(elems, element)
	}
	return
}

func (s1 *Set[T]) Union(s2 *Set[T]) (s *Set[T]) {
	s1.lock(); defer s1.unlock()
	s2.lock(); defer s2.unlock()

	s = NewSet[T]()
	for element := range s1.elements {
		s.Add(element)
	}

	for element := range s2.elements {
		s.Add(element)
	}
	return
}
