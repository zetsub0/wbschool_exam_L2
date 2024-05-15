package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
Со стороны кажется, что изменился класс объекта.

Плюсы:
	1 Избавляет от множества больших условных операторов машины состояний
	2 Упрощает код контекста.
	3 Концентрирует в одном месте код, связанный с определённым состоянием.

Минсуы:
	1 Может неоправданно усложнить код, если состояний мало и они редко меняются.

*/

type State interface {
	Handle(context *Context)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("State A handling request and changing state to B")
	context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("State B handling request and changing state to A")
	context.SetState(&ConcreteStateA{})
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	context := &Context{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
