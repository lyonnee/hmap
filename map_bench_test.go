package hmap

import (
	"testing"
)

func BenchmarkMap_NewMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMap[string, int](10)
	}
}

func BenchmarkMap_Len(b *testing.B) {
	m := NewMap[string, int](10)
	for i := 0; i < b.N; i++ {
		m.Len()
	}
}

func BenchmarkMap_Load(b *testing.B) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Load("key1")
	}
}

func BenchmarkMap_Swap(b *testing.B) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Swap("key1", 2)
	}
}

func BenchmarkMap_Store(b *testing.B) {
	m := NewMap[string, int](10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store("key1", 1)
	}
}

func BenchmarkMap_LoadOrStore(b *testing.B) {
	m := NewMap[string, int](10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.LoadOrStore("key1", 1)
	}
}

func BenchmarkMap_Delete(b *testing.B) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete("key1")
		m.Store("key1", 1)
	}
}

func BenchmarkMap_LoadAndDelete(b *testing.B) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.LoadAndDelete("key1")
		m.Store("key1", 1)
	}
}

func BenchmarkMap_Range(b *testing.B) {
	m := NewMap[string, int](10)
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

func BenchmarkMap_Clean(b *testing.B) {
	m := NewMap[string, int](10)
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

func BenchmarkMap_CompareAndSwap(b *testing.B) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.CompareAndSwap("key1", 1, 2)
	}
}

func BenchmarkMap_CompareAndDelete(b *testing.B) {
	m := NewMap[string, int](10)
	m.Store("key1", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.CompareAndDelete("key1", 1)
		m.Store("key1", 1)
	}
}
