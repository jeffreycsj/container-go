package cskiplist

import (
	"sync"
	"sync/atomic"
)

// ref: https://zhuanlan.zhihu.com/p/622177029

type node struct {
	key, val     any
	nexts        []*node
	sync.RWMutex // node mutex
}

// node mutex concurrent skip list
type ConcurrentSkipList struct {
	elementCnt atomic.Int32

	DeleteMutex sync.RWMutex

	// use in set
	keyMutexMap sync.Map

	head *node

	// reuse the node structure created and deleted in the skip list, reduce the gc pressure
	nodesCache sync.Pool

	compareFunc func(key1, key2 any) bool
}

// a concurrent safe skip list
func NewConcurrentSkipList(compareFunc func(key1, key2 any) bool) *ConcurrentSkipList {
	cskiplist := &ConcurrentSkipList{
		head: &node{
			nexts: make([]*node, 1),
		},
		nodesCache: sync.Pool{
			New: func() any {
				return &node{}
			},
		},
		compareFunc: compareFunc,
	}

	if cskiplist.compareFunc == nil {
		cskiplist.compareFunc = defaultCompareFunc
	}

	return cskiplist
}

func defaultCompareFunc(key1, key2 any) bool {
	return key1.(int) < key2.(int)
}

func (c *ConcurrentSkipList) Del(key any) {
	c.DeleteMutex.Lock()
	defer c.DeleteMutex.Unlock()

	var deleteNode *node

	move := c.head
	for level := len(c.head.nexts) - 1; level >= 0; level-- {
		for move.nexts[level] != nil && c.compareFunc(move.nexts[level].key, key) {
			move = move.nexts[level]
		}

		if move.nexts[level] == nil || (move.nexts[level].key != key && !c.compareFunc(move.nexts[level].key, key)) {
			continue
		}

		if deleteNode == nil {
			deleteNode = move.nexts[level]
		}

		move.nexts[level] = move.nexts[level].nexts[level]
	}

	if deleteNode == nil {
		return
	}

	defer c.elementCnt.Add(-1)
	deleteNode.nexts = nil
	c.nodesCache.Put(deleteNode)

	var dif int
	for level := len(c.head.nexts) - 1; level > 0; level-- {
		if c.head.nexts[level] != nil {
			break
		}
		dif++
	}
	c.head.nexts = c.head.nexts[:len(c.head.nexts)-dif]
}

func (c *ConcurrentSkipList) Get(key any) (any, bool) {
	c.DeleteMutex.RLock()
	defer c.DeleteMutex.RUnlock()
	if node_, exist := c.getNode(key); exist {
		return node_.val, true
	}
	return 0, false
}

func (c *ConcurrentSkipList) Set(key, val any) {
	c.DeleteMutex.RLock()
	defer c.DeleteMutex.RUnlock()

	keyMutex := c.getKeyMutex(key)
	keyMutex.Lock()
	defer keyMutex.Unlock()

	if node_, exist := c.getNode(key); exist {
		node_.val = val
		return
	}

	defer c.elementCnt.Add(1)

	newNode, _ := c.nodesCache.Get().(*node)
	newNode.key, newNode.val = key, val
	newNode.nexts = make([]*node, 2)

	newNode.Lock()
	defer newNode.Unlock()

	if 1 > len(c.head.nexts)-1 {
		c.head.Lock()
		for 1 > len(c.head.nexts)-1 {
			c.head.nexts = append(c.head.nexts, nil)
		}
		c.head.Unlock()
	}

	move := c.head
	var last *node
	for level := 1; level >= 0; level-- {
		for move.nexts[level] != nil && c.compareFunc(move.nexts[level].key, key) {
			move = move.nexts[level]
		}

		if move != last {
			move.Lock()
			defer move.Unlock()
			last = move
		}

		newNode.nexts[level] = move.nexts[level]
		move.nexts[level] = newNode
	}
}

func (c *ConcurrentSkipList) getKeyMutex(key any) *sync.Mutex {
	rawMutex, _ := c.keyMutexMap.LoadOrStore(key, &sync.Mutex{})
	mutex, _ := rawMutex.(*sync.Mutex)
	return mutex
}

func (c *ConcurrentSkipList) getNode(key any) (*node, bool) {
	move := c.head
	var last *node
	for level := len(c.head.nexts) - 1; level >= 0; level-- {
		for move.nexts[level] != nil && c.compareFunc(move.nexts[level].key, key) {
			move = move.nexts[level]
		}

		if move != last {
			move.RLock()
			defer move.RUnlock()
			last = move
		}

		if move.nexts[level] != nil && move.nexts[level].key == key {
			return move.nexts[level], true
		}
	}
	return nil, false
}
