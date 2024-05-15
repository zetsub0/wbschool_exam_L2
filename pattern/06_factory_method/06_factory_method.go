package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Паттерн определяет общий интерфейс для создания объектов в суперклассе.
Это позволяет подклассам изменять тип создаваемых объектов.

Плюсы:
	1 Выделяет код производства продуктов в одно место, упрощая поддержку кода.
 	2 Упрощает добавление новых продуктов в программу.
	3 Реализует принцип открытости/закрытости.

Минусы:
	1 Рискует стать супер-объектом,привазанным ко всем классам программы.


*/

type Transport interface {
	Deliver()
}

type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Доставка осуществляется грузовиком")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Доставка осуществляется кораблем")
}

type TransportFactory interface {
	CreateTransport() Transport
}

type TruckFactory struct{}

func (t *TruckFactory) CreateTransport() Transport {
	return &Truck{}
}

type ShipFactory struct{}

func (s *ShipFactory) CreateTransport() Transport {
	return &Ship{}
}

func main() {
	var factory TransportFactory

	factory = &TruckFactory{}
	truck := factory.CreateTransport()
	truck.Deliver()

	factory = &ShipFactory{}
	ship := factory.CreateTransport()
	ship.Deliver()
}
