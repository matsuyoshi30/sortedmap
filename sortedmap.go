package sortedmap

import (
	"cmp"
	"slices"
)

type SortedMap[K cmp.Ordered, V any] struct {
	kv    *m[K, V]
	sl    *[]K
	first *K
	last  *K
}

type m[K cmp.Ordered, V any] struct {
	kv map[K]V
}

func newM[K cmp.Ordered, V any]() *m[K, V] {
	return &m[K, V]{kv: make(map[K]V)}
}

func (m *m[K, V]) k(key K) (V, bool) {
	v, ok := m.kv[key]
	return v, ok
}

func (m *m[K, V]) p(key K, val V) {
	m.kv[key] = val
}

func (m *m[K, V]) d(key K) {
	delete(m.kv, key)
}

func (m *m[K, V]) len() int {
	return len(m.kv)
}

func NewSortedMap[K cmp.Ordered, V any]() *SortedMap[K, V] {
	sl := make([]K, 0)
	return &SortedMap[K, V]{kv: newM[K, V](), sl: &sl}
}

func (sm *SortedMap[K, V]) checkRange(key K) bool {
	if sm.first != nil && key < *sm.first {
		return false
	}
	if sm.last != nil && *sm.last < key { // TODO: exclude last
		return false
	}

	return true
}

func (sm *SortedMap[K, V]) Get(key K) (V, bool) {
	var ret V
	if !sm.checkRange(key) {
		return ret, false
	}

	v, ok := sm.kv.k(key)
	return v, ok
}

func (sm *SortedMap[K, V]) Put(key K, value V) V {
	var ret V
	v, ok := sm.kv.k(key)
	if ok {
		ret = v
	}
	sm.kv.p(key, value)
	sm.insert(key)
	return ret
}

func (sm *SortedMap[K, V]) insert(key K) {
	i, ok := slices.BinarySearch(*sm.sl, key)
	if ok {
		return
	}
	*sm.sl = slices.Insert(*sm.sl, i, key)
	if sm.first == nil || key < *sm.first {
		sm.first = &key
	}
	if sm.last == nil || *sm.last < key {
		sm.last = &key // TODO: last = key + 1
	}
}

func (sm *SortedMap[K, V]) Remove(key K) V {
	var ret V
	v, ok := sm.kv.k(key)
	if ok {
		ret = v
		sm.kv.d(key)
		sm.delete(key)
	}
	return ret
}

func (sm *SortedMap[K, V]) delete(key K) {
	i, ok := slices.BinarySearch(*sm.sl, key)
	if !ok {
		panic("not found")
	}
	sl := *sm.sl
	copy(sl[i:], sl[i+1:])
	*sm.sl = sl[:len(sl)-1]
}

func (sm *SortedMap[K, V]) FirstKey() (K, bool) {
	if sm.first == nil {
		var ret K
		return ret, false
	}
	return *sm.first, true
}

func (sm *SortedMap[K, V]) LastKey() (K, bool) {
	if sm.last == nil {
		var ret K
		return ret, false
	}
	return *sm.last, true
}

func (sm *SortedMap[K, V]) SubMap(from, to K) *SortedMap[K, V] {
	if sm.first == nil || from < *sm.first {
		return nil
	}
	if sm.last == nil || *sm.last < to {
		return nil
	}

	sm2 := NewSortedMap[K, V]()
	sm2.kv = sm.kv
	sm2.sl = sm.sl
	sm2.first = &from
	sm2.last = &to
	return sm2
}

func (sm *SortedMap[K, V]) HeadMap(to K) *SortedMap[K, V] {
	if sm.first == nil || to < *sm.first {
		return nil
	}

	sm2 := NewSortedMap[K, V]()
	sm2.kv = sm.kv
	sm2.sl = sm.sl
	sm2.last = &to
	return sm2
}

func (sm *SortedMap[K, V]) TailMap(from K) *SortedMap[K, V] {
	if sm.last == nil || *sm.last < from {
		return nil
	}

	sm2 := NewSortedMap[K, V]()
	sm2.kv = sm.kv
	sm2.sl = sm.sl
	sm2.first = &from

	return sm2
}

func (sm *SortedMap[K, V]) IsEmpty() bool {
	return sm.kv.len() == 0 || *sm.last <= *sm.first
}
