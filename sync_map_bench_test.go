package hmap

import (
	"testing"
)

func BenchmarkSyncMap_NewSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSyncMap[string, int]()
	}
}

func BenchmarkSyncMap_Len(b *testing.B) {
	m := NewSyncMap[string, int]()
	for i := 0; i < b.N; i++ {
		m.Len()
	}
}

func BenchmarkSyncMap_Load(b *testing.B) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Load("key1")
	}
}

func BenchmarkSyncMap_Swap(b *testing.B) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Swap("key1", 2)
	}
}

func BenchmarkSyncMap_Store(b *testing.B) {
	m := NewSyncMap[string, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store("key1", 1)
	}
}

func BenchmarkSyncMap_LoadOrStore(b *testing.B) {
	m := NewSyncMap[string, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.LoadOrStore("key1", 1)
	}
}

func BenchmarkSyncMap_Delete(b *testing.B) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete("key1")
		m.Store("key1", 1)
	}
}

func BenchmarkSyncMap_LoadAndDelete(b *testing.B) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.LoadAndDelete("key1")
		m.Store("key1", 1)
	}
}

func BenchmarkSyncMap_Range(b *testing.B) {
	m := NewSyncMap[string, int]()
	for i := 0; i < 100; i++ {
		m.Store("key"+string(rune(i)), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Range(func(k string, v int) bool {
			return true
		})
	}
}

func BenchmarkSyncMap_Clean(b *testing.B) {
	m := NewSyncMap[string, int]()
	for i := 0; i < 100; i++ {
		m.Store("key"+string(rune(i)), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Clean()
		for j := 0; j < 100; j++ {
			m.Store("key"+string(rune(j)), j)
		}
	}
}

func BenchmarkSyncMap_CompareAndSwap(b *testing.B) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.CompareAndSwap("key1", 1, 2)
	}
}

func BenchmarkSyncMap_CompareAndDelete(b *testing.B) {
	m := NewSyncMap[string, int]()
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.CompareAndDelete("key1", 1)
		m.Store("key1", 1)
	}
}
