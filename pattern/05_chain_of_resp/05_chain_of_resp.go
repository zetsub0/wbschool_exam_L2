package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Паттерн позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

Плюсы:
	1 Уменьшает зависимость между клиентом и обработчиками.
 	2 Реализует принцип единственной обязанности.
	3 Реализует принцип открытости/закрытости.

Минусы:
	1 Нет гарантии обработки запроса.

*/

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled request")
	} else {
		fmt.Println("ConcreteHandlerA passed request")
		h.BaseHandler.Handle(request)
	}
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled request")
	} else {
		fmt.Println("ConcreteHandlerB passed request")
		h.BaseHandler.Handle(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	fmt.Println("Sending request 'A':")
	handlerA.Handle("A")

	fmt.Println("\nSending request 'B':")
	handlerA.Handle("B")

	fmt.Println("\nSending request 'C':")
	handlerA.Handle("C")
}
