package hmap

import (
	"sync"
	"sync/atomic"
)

type SyncMap[K comparable, V any] struct {
	d   *sync.Map
	len atomic.Int64
}

func NewSyncMap[K comparable, V any]() HMap[K, V] {
	return &SyncMap[K, V]{
		d:   new(sync.Map),
		len: atomic.Int64{},
	}
}

func (m *SyncMap[K, V]) Len() int {
	return int(m.len.Load())
}

func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	var ret V
	if v, ok := m.d.Load(key); ok {
		if v != nil {
			ret = v.(V)
		}
		return ret, true
	}

	// not found
	return ret, false
}

func (m *SyncMap[K, V]) Swap(k K, v V) (V, bool) {
	var ret = v
	previous, loaded := m.d.Swap(k, v)
	if !loaded {
		// not exist, len +1
		m.len.Add(1)
	} else {
		if previous != nil {
			ret = previous.(V)
		}
	}

	return ret, loaded
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	_, loaded := m.d.Swap(k, v)
	if !loaded {
		// not exist, len +1
		m.len.Add(1)
	}
}

func (m *SyncMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	d, ok := m.d.LoadOrStore(k, v)
	if ok {
		return d.(V), ok
	}

	// not exist, len +1
	m.len.Add(1)

	return v, ok
}

func (m *SyncMap[K, V]) Delete(k K) {
	_, ok := m.d.LoadAndDelete(k)
	if ok {
		// exist, len -1
		m.len.Add(-1)
	}
}

func (m *SyncMap[K, V]) LoadAndDelete(k K) (V, bool) {
	var ret V
	v, ok := m.d.LoadAndDelete(k)
	if ok {
		// exist, len -1
		m.len.Add(-1)
		if v != nil {
			ret = v.(V)
		}
	}
	return ret, ok
}

func (m *SyncMap[K, V]) Range(fn func(k K, v V) bool) {
	m.d.Range(func(key, value any) bool {
		k := key.(K)
		v := value.(V)

		if !fn(k, v) {
			return false
		}

		return true
	})
}

func (m *SyncMap[K, V]) Clean() {
	m.d = new(sync.Map)
	m.len = atomic.Int64{}
}

func (m *SyncMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.d.CompareAndSwap(key, old, new)
}

func (m *SyncMap[K, V]) CompareAndDelete(key K, old V) bool {
	del := m.d.CompareAndDelete(key, old)
	if del {
		// deleted, len -1
		m.len.Add(-1)
	}
	return del
}
