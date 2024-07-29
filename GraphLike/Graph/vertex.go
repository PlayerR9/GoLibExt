package Graph

import (
	uc "github.com/PlayerR9/lib_units/common"
)

type VertexElementer interface {
	uc.Equaler
	uc.Copier
}

type Vertex[T VertexElementer] struct {
	value T

	isInitial bool
	isFinal   bool
}

func (v *Vertex[T]) Equals(other uc.Equaler) bool {
	if other == nil {
		return false
	}

	otherV, ok := other.(*Vertex[T])
	if !ok {
		return false
	}

	return v.value.Equals(otherV.value) &&
		v.isInitial == otherV.isInitial &&
		v.isFinal == otherV.isFinal
}

func (v *Vertex[T]) Copy() uc.Copier {
	return &Vertex[T]{
		value:     v.value.Copy().(T),
		isInitial: v.isInitial,
		isFinal:   v.isFinal,
	}
}

func NewVertex[T VertexElementer](value T) *Vertex[T] {
	return &Vertex[T]{
		value:     value,
		isInitial: false,
		isFinal:   false,
	}
}

func (v *Vertex[T]) GetValue() T {
	return v.value
}
