package hmap

import (
	"testing"
)

func TestSyncMap_NewSyncMap(t *testing.T) {
	m := NewSyncMap[string, int]()
	if m == nil {
		t.Errorf("NewSyncMap returned nil")
	}
}

func TestSyncMap_Len(t *testing.T) {
	m := NewSyncMap[string, int]()
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
	m.Store("key1", 1)
	if m.Len() != 1 {
		t.Errorf("Expected length 1, got %d", m.Len())
	}
}

func TestSyncMap_Load(t *testing.T) {
	m := NewSyncMap[string, int]()
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

func TestSyncMap_Swap(t *testing.T) {
	m := NewSyncMap[string, int]()
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

func TestSyncMap_Store(t *testing.T) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	v, ok := m.Load("key1")
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	m := NewSyncMap[string, int]()
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

func TestSyncMap_Delete(t *testing.T) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	m.Delete("key1")
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}

func TestSyncMap_LoadAndDelete(t *testing.T) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	v, ok := m.LoadAndDelete("key1")
	if !ok || v != 1 {
		t.Errorf("Expected value 1, got %d, ok %v", v, ok)
	}
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}

func TestSyncMap_Range(t *testing.T) {
	m := NewSyncMap[string, int]()
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

func TestSyncMap_Clean(t *testing.T) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	m.Clean()
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}

func TestSyncMap_CompareAndSwap(t *testing.T) {
	m := NewSyncMap[string, int]()
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

func TestSyncMap_CompareAndDelete(t *testing.T) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	deleted := m.CompareAndDelete("key1", 1)
	if !deleted {
		t.Errorf("Expected deleted to be true, got %v", deleted)
	}
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}
}
