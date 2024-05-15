package main

import "fmt"

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Паттерн превращает запросы в объекты.
Это позволяет передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их.

Плюсы:
	1 Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	2 Позволяет реализовать отложенный запуск операций.
	3 Реализует принцип открытости/закрытости.

Минусы:
	1 Много доп класов.

*/

type Command interface {
	Execute()
}

type Order struct {
	details string
}

type Waiter struct {
	orders []Command
}

func (w *Waiter) TakeOrder(command Command) {
	w.orders = append(w.orders, command)
	fmt.Println("Официант принял заказ")
}

func (w *Waiter) PlaceOrders() {
	for _, command := range w.orders {
		command.Execute()
	}
}

type Chef struct{}

func (c *Chef) PrepareOrder(orderDetails string) {
	fmt.Printf("Повар готовит: %s\n", orderDetails)
}

type FoodOrder struct {
	chef  *Chef
	order *Order
}

func (f *FoodOrder) Execute() {
	f.chef.PrepareOrder(f.order.details)
}

func main() {
	chef := &Chef{}

	order1 := &Order{details: "Салат Цезарь"}
	order2 := &Order{details: "Стейк с кровью"}

	foodOrder1 := &FoodOrder{chef: chef, order: order1}
	foodOrder2 := &FoodOrder{chef: chef, order: order2}

	waiter := &Waiter{}
	waiter.TakeOrder(foodOrder1)
	waiter.TakeOrder(foodOrder2)

	waiter.PlaceOrders()
}
