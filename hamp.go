package hmap

import "sync"

type HMap[K, V any] struct {
	d *sync.Map
	l int
}

func New[K, V any]() *HMap[K, V] {
	return &HMap[K, V]{
		d: new(sync.Map),
		l: 0,
	}
}

func (m *HMap[K, V]) Len() int {
	return m.l
}

func (m *HMap[K, V]) Load(key K) (V, bool) {
	if v, ok := m.d.Load(key); ok {
		return v.(V), true
	}

	// not found
	var v V
	return v, false
}

func (m *HMap[K, V]) Swap(k K, v V) (V, bool) {
	previous, loaded := m.d.Swap(k, v)
	if !loaded {
		// not exist, len ++
		m.l++
	}
	return previous.(V), loaded
}

func (m *HMap[K, V]) Store(k K, v V) {
	_, loaded := m.d.Swap(k, v)
	if !loaded {
		// not exist, len ++
		m.l++
	}
}

func (m *HMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	res, ok := m.d.LoadOrStore(k, v)
	if !ok {
		// not exist, len ++
		m.l++
	}

	return res.(V), ok
}

func (m *HMap[K, V]) Delete(k K) {
	_, ok := m.d.LoadAndDelete(k)
	if ok {
		// exist, len --
		m.l--
	}
}

func (m *HMap[K, V]) LoadAndDelete(k K) (V, bool) {
	v, ok := m.d.LoadAndDelete(k)
	if ok {
		// exist, len --
		m.l--
	}

	return v.(V), ok
}

func (m *HMap[K, V]) Range(fn func(k any, v any) bool) {
	m.d.Range(fn)
}

func (m *HMap[K, V]) Clean() {
	m.d = new(sync.Map)
	m.l = 0
}
