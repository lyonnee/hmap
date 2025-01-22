package hmap

import (
	"sync"
)

type Map[K comparable, V any] struct {
	mutex sync.Mutex
	d     map[K]V
}

func NewMap[K comparable, V any](size int) HMap[K, V] {
	return &Map[K, V]{
		mutex: sync.Mutex{},
		d:     make(map[K]V, size),
	}
}

func (m *Map[K, V]) Len() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return len(m.d)
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if v, ok := m.d[key]; ok {
		return v, true
	}

	// not found
	var ret V
	return ret, false
}

func (m *Map[K, V]) Swap(k K, v V) (V, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	value, ok := m.d[k]
	m.d[k] = v

	return value, ok
}

func (m *Map[K, V]) Store(k K, v V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.d[k] = v
}

func (m *Map[K, V]) LoadOrStore(k K, v V) (V, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if ret, ok := m.d[k]; ok {
		return ret, ok
	}

	m.d[k] = v

	var ret V
	return ret, false
}

func (m *Map[K, V]) Delete(k K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.d, k)
}

func (m *Map[K, V]) LoadAndDelete(k K) (V, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, ok := m.d[k]
	if ok {
		delete(m.d, k)
	}

	return v, ok
}

func (m *Map[K, V]) Range(fn func(k K, v V) bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for k, v := range m.d {
		if !fn(k, v) {
			return
		}
	}
}

func (m *Map[K, V]) Clean() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.d = make(map[K]V)
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, ok := m.d[key]
	if ok && any(v) == any(old) {
		m.d[key] = new
		return true
	}

	return false
}

func (m *Map[K, V]) CompareAndDelete(key K, old V) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	v, ok := m.d[key]
	if ok && any(v) == any(old) {
		delete(m.d, key)
		return true
	}

	return false
}
