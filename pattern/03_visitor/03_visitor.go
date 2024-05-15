package main

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern

Паттерн позволяет добавлять в программу новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.

Плюсы:
	1 Упрощает добавление операций, работающих со сложными структурами объектов.
	2 Объединяет родственные операции в одном классе.

Минусы:
	1 Может привести к нарушению инкапсуляции элементов.
	2 Tе оправдан, если иерархия компонентов часто меняется

*/

import "fmt"

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct {
	name string
}

func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

type ConcreteElementB struct {
	value int
}

func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("ConcreteVisitor1: Visiting ConcreteElementA with name: %s\n", e.name)
}

func (v *ConcreteVisitor1) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("ConcreteVisitor1: Visiting ConcreteElementB with value: %d\n", e.value)
}

type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("ConcreteVisitor2: Visiting ConcreteElementA with name: %s\n", e.name)
}

func (v *ConcreteVisitor2) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("ConcreteVisitor2: Visiting ConcreteElementB with value: %d\n", e.value)
}

func main() {
	elements := []Element{
		&ConcreteElementA{name: "Element A1"},
		&ConcreteElementB{value: 42},
	}

	visitors := []Visitor{
		&ConcreteVisitor1{},
		&ConcreteVisitor2{},
	}

	for _, element := range elements {
		for _, visitor := range visitors {
			element.Accept(visitor)
		}
	}
}
