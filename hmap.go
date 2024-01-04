package hmap

import (
	"sync"
	"sync/atomic"
)

type Map[K comparable, V any] struct {
	d *sync.Map
	l atomic.Int64
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		d: new(sync.Map),
		l: atomic.Int64{},
	}
}

func (m *Map[K, V]) Len() int {
	return int(m.l.Load())
}

func (m *Map[K, V]) Load(key K) (V, bool) {
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

func (m *Map[K, V]) Swap(k K, v V) (V, bool) {
	var ret V
	previous, loaded := m.d.Swap(k, v)
	if !loaded {
		// not exist, len +1
		m.l.Add(1)
		ret = v
	} else {
		if previous != nil {
			ret = previous.(V)
		}
	}
	return ret, loaded
}

func (m *Map[K, V]) Store(k K, v V) {
	_, loaded := m.d.Swap(k, v)
	if !loaded {
		// not exist, len +1
		m.l.Add(1)
	}
}

func (m *Map[K, V]) LoadOrStore(k K, v V) (V, bool) {
	var ret V
	d, ok := m.d.LoadOrStore(k, v)
	if !ok {
		if d != nil {
			ret = d.(V)
		}
		// not exist, len +1
		m.l.Add(1)
	}

	ret = v

	return ret, ok
}

func (m *Map[K, V]) Delete(k K) {
	_, ok := m.d.LoadAndDelete(k)
	if ok {
		// exist, len -1
		m.l.Add(-1)
	}
}

func (m *Map[K, V]) LoadAndDelete(k K) (V, bool) {
	var ret V
	v, ok := m.d.LoadAndDelete(k)
	if ok {
		// exist, len -1
		m.l.Add(-1)
		if v != nil {
			ret = v.(V)
		}
	}
	return ret, ok
}

func (m *Map[K, V]) Range(fn func(k any, v any) bool) {
	m.d.Range(fn)
}

func (m *Map[K, V]) Clean() {
	m.d = new(sync.Map)
	m.l = atomic.Int64{}
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.d.CompareAndSwap(key, old, new)
}

func (m *Map[K, V]) CompareAndDelete(key K, old V) bool {
	del := m.d.CompareAndDelete(key, old)
	if del {
		// deleted, len -1
		m.l.Add(-1)
	}
	return del
}
