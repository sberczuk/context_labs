package pkg

import "cmp"

type AttribName string

// List attribute names here
const Color = AttribName("COLOR")
const Price = AttribName("PRICE")
const Quantity = AttribName("QUANTITY")

type Attribute interface {
	GetName() AttribName
	GetValue() any
}

type BaseAttribute[T cmp.Ordered] struct {
	Name  AttribName `json:"name,omitempty"`
	Value T          `json:"value,omitempty"`
}

func (a BaseAttribute[T]) GetName() AttribName {
	return a.Name
}

func (a BaseAttribute[T]) Equals(v any) bool {
	return a.Value == v
}

func (a BaseAttribute[T]) LessThan(v any) bool {
	// This type conversion is potentially unsafe, I could change this to return an error,
	// I could keep a name/type mapping somewhere.
	// this comment applies to all of these methods
	return a.Value < v.(T)
}

func (a BaseAttribute[T]) GreaterThan(v any) bool {
	// This type conversion is potentially unsafe, I could change this to return an error,
	// I could keep a name/type mapping somewhere.
	// this comment applies to all of these methods
	return a.Value > v.(T)
}

// GetValue returns the value. Type casting is handled by the caller.
func (a BaseAttribute[T]) GetValue() any {
	return a.Value
}
