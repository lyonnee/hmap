package hmap

import (
	"testing"
)

func TestMap_NewMap(t *testing.T) {
	m := NewMap[string, int](10)
	if m == nil {
		t.Errorf("NewMap returned nil")
	}
}

func TestMap_Len(t *testing.T) {
	m := NewMap[string, int](10)
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
	m.Store("key1", 1)
	if m.Len() != 1 {
		t.Errorf("Expected length 1, got %d", m.Len())
	}
}

func TestMap_Load(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	v, ok := m.Load("key1")
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
	v, ok = m.Load("key2")
	if ok {
		t.Errorf("Expected ok to be false, got %v", ok)
	}
}

func TestMap_Swap(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	v, ok := m.Swap("key1", 2)
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
	if m.Len() != 1 {
		t.Errorf("Expected length 1, got %d", m.Len())
	}
	v, ok = m.Swap("key2", 2)
	if ok {
		t.Errorf("Expected ok to be false, got %v", ok)
	}
	if m.Len() != 2 {
		t.Errorf("Expected length 2, got %d", m.Len())
	}
}

func TestMap_Store(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	v, ok := m.Load("key1")
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
}

func TestMap_LoadOrStore(t *testing.T) {
	m := NewMap[string, int](10)
	v, ok := m.LoadOrStore("key1", 1)
	if ok {
		t.Errorf("Expected ok to be false, got %v", ok)
	}
	if m.Len() != 1 {
		t.Errorf("Expected length 1, got %d", m.Len())
	}
	v, ok = m.LoadOrStore("key1", 2)
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
}

func TestMap_Delete(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	m.Delete("key1")
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}

func TestMap_LoadAndDelete(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	v, ok := m.LoadAndDelete("key1")
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}

func TestMap_Range(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	m.Store("key2", 2)
	count := 0
	m.Range(func(k string, v int) bool {
		count++
		return true
	})
	if count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestMap_Clean(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	m.Clean()
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}

func TestMap_CompareAndSwap(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	swapped := m.CompareAndSwap("key1", 1, 2)
	if !swapped {
		t.Errorf("Expected swapped to be true, got %v", swapped)
	}
	v, ok := m.Load("key1")
	if !ok || v != 2 {
		t.Errorf("Expected value 2, got %d, ok %v", v, ok)
	}
}

func TestMap_CompareAndDelete(t *testing.T) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	deleted := m.CompareAndDelete("key1", 1)
	if !deleted {
		t.Errorf("Expected deleted to be true, got %v", deleted)
	}
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}
