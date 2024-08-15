package set

type SetKey interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

// Set is a type alias for a map that stores the set elements.
type Set [T SetKey] struct {
	m map[T]struct{}
}

func NewSet[T SetKey]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// Add adds an element to the set.
func (s *Set[T]) Add(element T) {
    s.m[element] = struct{}{}
}

// Remove removes an element from the set.
func (s *Set[T]) Remove(element T) {
    delete(s.m, element)
}

// Contains checks if an element is in the set.
func (s *Set[T]) Contains(element T) bool {
    _, exists := s.m[element]
    return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
    return len(s.m)
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
    for key := range s.m {
        delete(s.m, key)
    }
}