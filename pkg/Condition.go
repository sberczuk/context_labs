package pkg

import "cmp"

type Condition interface {

	// Apply the condition to the product
	Apply(p *Product) bool
}

// Use Generics to support different types
type ConcreteCondition[T cmp.Ordered] struct {
	AttributeName AttribName
	Operator      Operator[T]
	Value         T
}

func NewConcreteCondition[T cmp.Ordered](attributeName AttribName, operator Operator[T], value T) *ConcreteCondition[T] {
	return &ConcreteCondition[T]{
		AttributeName: attributeName,
		Value:         value,
		Operator:      operator,
	}
}

func (c *ConcreteCondition[T]) Apply(p *Product) bool {

	//TODO: DO Type Switch here?
	attribute := p.GetAttribute(c.AttributeName)
	value := attribute.GetValue().(T)
	return c.Operator.Apply(c.Value, value)
}

type Operator[T cmp.Ordered] interface {
	Apply(l, r T) bool
}

type EqualOperator[T cmp.Ordered] struct {
}

func (e *EqualOperator[T]) Apply(l, r T) bool {
	return l == r
}

type LessThanOperator[T cmp.Ordered] struct {
}

func (e *LessThanOperator[T]) Apply(l, r T) bool {
	return l < r
}

type GreaterThanOperator[T cmp.Ordered] struct {
}

func (e *GreaterThanOperator[T]) Apply(l, r T) bool {
	return l > r
}
