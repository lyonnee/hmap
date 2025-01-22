<div align="center">
</br>

# HMap - Safe For Concurrent Generic Map

| English | [中文](README_zh.md) |
| --- | --- |

HMap is a safe for concurrent generic map implementation in Go that provides a simple and efficient API. It supports generic key-value pairs and includes methods for querying the length of the map as well as clearing the entire map.

</div>

## Features

- **Generic Support:** Utilizes Go's generic features for flexibility with different types of key-value pairs.
- **Length Querying:** Easily determine the length of the map using the `Len` method.
- **Clear Map:** Clear the entire map using the `Clean` method.

## Go Version

`>= 1.20`

## Quick Start

### Using `Map`

`Map` is a concurrent-safe map implementation based on a mutex.

```go
package main

import (
	"fmt"
	"github.com/lyonnee/hmap"
)

func main() {
	// Create a new Map instance
	myMap := hmap.NewMap[string, int](10)

	// Store key-value pairs
	myMap.Store("key1", 1)
	myMap.Store("key2", 2)

	// Query the length of the map
	length := myMap.Len()
	fmt.Printf("Map Length: %d\n", length)

	// Clear the map
	myMap.Clean()

	// Query the length of the map after cleaning
	length = myMap.Len()
	fmt.Printf("Map Length after Clean: %d\n", length)
}
```

### Using `SyncMap`

`SyncMap` is a concurrent-safe map implementation based on sync.Map.

```go
package main

import (
	"fmt"
	"github.com/lyonnee/hmap"
)

func main() {
	// Create a new SyncMap instance
	mySyncMap := hmap.NewSyncMap[string, int]()

	// Store key-value pairs
	mySyncMap.Store("key1", 1)
	mySyncMap.Store("key2", 2)

	// Query the length of the map
	length := mySyncMap.Len()
	fmt.Printf("SyncMap Length: %d\n", length)

	// Clear the map
	mySyncMap.Clean()

	// Query the length of the map after cleaning
	length = mySyncMap.Len()
	fmt.Printf("SyncMap Length after Clean: %d\n", length)
}
```

## Comparison of `Map` and `SyncMap`
### Map
- Advantages:
	- mplemented using a mutex, simple and intuitive.
	- Suitable for scenarios with low concurrency read/write operations.
- Disadvantages:
	- Performance may be lower than `SyncMap` in high-concurrency write scenarios, as the mutex causes write operations to block.
### SyncMap
- Advantages:
	- Implemented using sync.Map, suitable for high-concurrency scenarios.
	- Write operations do not block read operations, providing better performance.
- Disadvantages:
	- Implementation is more complex and has a larger codebase.
	- Performance may be lower than `Map` in low-concurrency scenarios.

## Methods

### `New() *Map[K, V]`

Creates a new instance of the HMap.

### `Len() int`

Gets the length of the map.

### `Load(key K) (V, bool)`

Loads the value associated with the specified key.

### `Swap(k K, v V) (V, bool)`

Swaps the value associated with the specified key and returns the previous value. If the key does not exist, the length is incremented.

### `Store(k K, v V)`

Stores the key-value pair. If the key already exists, its value is updated.

### `LoadOrStore(k K, v V) (V, bool)`

Loads the value for the specified key. If the key does not exist, the given value is stored, and the length is incremented.

### `Delete(k K)`

Deletes the value associated with the specified key. If the key exists, the length is decremented.

### `LoadAndDelete(k K) (V, bool)`

Loads and deletes the value associated with the specified key. If the key exists, the length is decremented.

### `Range(fn func(k K, v V) bool)`

Iterates over the map, applying a given function to each key-value pair.

### `Clean()`

Clears the entire map.

### `CompareAndSwap(key K, old, new V) bool`

Compares and swaps the value associated with the specified key. Returns `true` if the swap is successful.

### `CompareAndDelete(key K, old V) bool`

Compares and deletes the value associated with the specified key, given the old value. Returns `true` if the delete is successful. If successful, the length is decremented.

## License

This project is distributed under the [MIT License](LICENSE).

## Contributors

Thanks to all contributors to this project.

## Acknowledgments

Thank you for using HMap! If you have any questions or suggestions, feel free to reach out.