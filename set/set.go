package set

import "sync"

var pool = sync.Pool{}

// Set is a type alias for a map that stores the set elements.
type Set struct {
	items map[interface{}]struct{}
    lock sync.RWMutex
}

// Add adds an element to the set.
func (s *Set) Add(element interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.items[element] = struct{}{}
}

// Remove removes an element from the set.
func (s *Set) Remove(element interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()
    delete(s.items, element)
}

// Contains checks if an element is in the set.
func (s *Set) Contains(element interface{}) bool {
    s.lock.RLock()
    defer s.lock.RUnlock()
    _, exists := s.items[element]
    return exists
}

// Size returns the number of elements in the set.
func (s *Set) Size() int {
    s.lock.RLock()
    defer s.lock.RUnlock()
    return len(s.items)
}

// Clear removes all elements from the set.
func (s *Set) Clear() {
    s.lock.Lock()
    defer s.lock.Unlock()
    for key := range s.items {
        delete(s.items, key)
    }
}

func (s *Set) Items() []interface{} {
    s.lock.RLock()
    defer s.lock.RUnlock()
    items := make([]interface{}, 0, len(s.items))
    for item := range s.items {
        items = append(items, item)
    }

    return items
}

func NewSet(items ...interface{}) *Set {
	set := pool.Get().(*Set)
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	return set
}

func init() {
	pool.New = func() interface{} {
		return &Set{
			items: make(map[interface{}]struct{}, 10),
		}
	}
}