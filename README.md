# Queue

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/goneric/queue)
[![Build Status](https://github.com/goneric/queue/actions/workflows/build.yaml/badge.svg)](https://github.com/goneric/queue/actions)
[![codecov](https://codecov.io/gh/goneric/queue/branch/main/graph/badge.svg)](https://codecov.io/gh/goneric/queue)

A lock-free queue implement in Go using generics.

## 🎯 Features
- ✅ Simple, lightweight without any external dependencies
- ✅ Type-safe with Go generics (required Go 1.18)
- ✅ Use with any Go data types, `bool`, `int`, `string`, structs or pointers
  
## ⚡️ Usage
```go
import "github.com/goneric/queue"
```

```go
type Queue[V any] interface {
	Push(value V)
	Pop() (value V, ok bool)
	Peek() (value V, ok bool)
	Len() int
}
```

### Examples