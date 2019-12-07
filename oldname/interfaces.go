package oldname

// Value will tell you if it is cold.
type Value interface {
	IsCold() bool
}

// Container contains some values
type Container interface {
	WithHandle(handle int) Container
	Values() []Value
}
