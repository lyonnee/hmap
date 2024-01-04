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

```go
import (
	"fmt"
	"github.com/lyonnee/hmap"
)

func main() {
	// 创建一个新的 HMap
	myMap := hmap.New[int, string]()

	// 存储键值对
	myMap.Store(1, "One")
	myMap.Store(2, "Two")

	// 查询映射长度
	length := myMap.Len()
	fmt.Printf("Map Length: %d\n", length)

	// 清空映射
	myMap.Clean()

	// 再次查询映射长度
	length = myMap.Len()
	fmt.Printf("Map Length after Clean: %d\n", length)
}
```

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

### `Range(fn func(k any, v any) bool)`

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