<div align="center">
</br>

# HMap - 并发安全的泛型Map

| [English](README.md) | 中文 |
| --- | --- |

HMap 是一个并发安全的支持泛型的Map实现，旨在提供简单易用的 API，支持泛型键值对，同时添加了对映射长度的查询以及清空映射的方法。

</div>

## 特性

- **泛型支持:** 使用 Golang 的泛型特性，可以在不同类型的键值对上使用 HMap。
- **映射长度查询:** 通过 `Len` 方法快速获取映射的长度。
- **清空映射:** 使用 `Clean` 方法轻松清空整个映射。

## Go版本

`>= 1.20`

## 快速开始

### 使用 `Map`

`Map` 是一个基于互斥锁的并发安全 Map 实现。

```go
package main

import (
	"fmt"
	"github.com/lyonnee/hmap"
)

func main() {
	// 创建一个新的 Map 实例
	myMap := hmap.NewMap[string, int](10)

	// 存储键值对
	myMap.Store("key1", 1)
	myMap.Store("key2", 2)

	// 查询 Map 的长度
	length := myMap.Len()
	fmt.Printf("Map Length: %d\n", length)

	// 清空 Map
	myMap.Clean()

	// 查询清空后的 Map 长度
	length = myMap.Len()
	fmt.Printf("Map Length after Clean: %d\n", length)
}
```

### 使用 `SyncMap`

`SyncMap` 是一个基于 sync.Map 的并发安全 Map 实现。

```go
package main

import (
	"fmt"
	"github.com/lyonnee/hmap"
)

func main() {
	// 创建一个新的 SyncMap 实例
	mySyncMap := hmap.NewSyncMap[string, int]()

	// 存储键值对
	mySyncMap.Store("key1", 1)
	mySyncMap.Store("key2", 2)

	// 查询 Map 的长度
	length := mySyncMap.Len()
	fmt.Printf("SyncMap Length: %d\n", length)

	// 清空 Map
	mySyncMap.Clean()

	// 查询清空后的 Map 长度
	length = mySyncMap.Len()
	fmt.Printf("SyncMap Length after Clean: %d\n", length)
}
```

## `Map` 和 `SyncMap` 的比较
### Map
- 优点：
	- 使用互斥锁实现，简单直观。
	- 适用于并发读写较少的场景。
- 缺点：
	-在高并发写入时，性能可能不如 SyncMap，因为互斥锁会导致写操作阻塞。
### SyncMap
- 优点：
	- 使用 sync.Map 实现，适合高并发场景。
	- 写操作不会阻塞读操作，性能更高。
- 缺点：
	- 实现相对复杂，代码量较大。
	- 在低并发场景下，性能可能不如 Map。

## 方法

### `New() *HMap[K, V]`

创建一个新的 HMap 实例。

### `Len() int`

获取映射的长度。

### `Load(key K) (V, bool)`

加载指定键的值。

### `Swap(k K, v V) (V, bool)`

交换指定键的值，并返回之前的值。

### `Store(k K, v V)`

存储键值对。如果键已存在，则更新其值。

### `LoadOrStore(k K, v V) (V, bool)`

加载键的值，如果键不存在，则存储给定的值。

### `Delete(k K)`

删除指定键的值。

### `LoadAndDelete(k K) (V, bool)`

加载并删除指定键的值。

### `Range(fn func(k K, v V) bool)`

遍历映射，对每个键值对执行指定的函数。

### `Clean()`

清空整个映射。

### `CompareAndSwap(key K, old, new V) bool`

比较并替换与指定键关联的值。如果替换成功，则返回“true”。

### `CompareAndDelete(key K, old V) bool`

给定旧值，比较并删除与指定键关联的值。如果删除成功，则返回“true”。如果成功，则长度递减。

## 许可证

该项目基于 [MIT 许可证](LICENSE) 进行分发。

## 贡献者

感谢所有为该项目做出贡献的人。

## 致谢

感谢您使用 HMap ！如有问题或建议，请随时提出。