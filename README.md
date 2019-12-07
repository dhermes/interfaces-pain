# Package Renames and Interfaces Don't Mix

> Pain Caused When Migrating a Golang Interface during a package rename

When moving the interfaces

```go
type Value interface {
	IsCold() bool
}

type Container interface {
	WithHandle(handle int) Container
	Values() []Value
}
```

from one package to another (during a rename), it is not possible for a
value to satisfy both interfaces. For example, the function

```go
func UpgradeContainer(c oldname.Container) newname.Container {
	return c
}
```

does not compile due to

```
$ go build ./...
# github.com/dhermes/interfaces-pain/consumer
consumer/usage.go:15:2: cannot use c (type oldname.Container) as type newname.Container in return argument:
        oldname.Container does not implement newname.Container (wrong type for Values method)
                have Values() []oldname.Value
                want Values() []newname.Value
```

**in spite** of the fact that `oldname.Value` and `newname.Value` are the
same interface as verified by the fact that

```go
func UpgradeValue(v oldname.Value) newname.Value {
	return v
}
```

does not produce compiler errors.
