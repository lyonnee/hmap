package hmap

type HMap[K comparable, V any] interface {
	Len() int
	Load(key K) (V, bool)
	Swap(k K, v V) (V, bool)
	Store(k K, v V)
	LoadOrStore(k K, v V) (V, bool)
	Delete(k K)
	LoadAndDelete(k K) (V, bool)
	Range(fn func(k K, v V) bool)
	Clean()
	CompareAndSwap(key K, old, new V) bool
	CompareAndDelete(key K, old V) bool
}
