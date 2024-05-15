package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Паттерн определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс.
После этого алгоритмы можно взаимозаменять прямо во время исполнения программы.

Плюсы:
	1 Изолирует код и данные алгоритмов от остальных классов.
	2 Горячая замена алгоритмов на лету.
	3 Реализует принцип открытости/закрытости.

Минусы:
	1 Усложняет программу за счёт дополнительных классов.
	2 Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

*/

type RouteStrategy interface {
	CalculateRoute(start, end string) string
}

type CarRouteStrategy struct{}

func (c *CarRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Расчет маршрута на автомобиле из %s в %s", start, end)
}

type WalkRouteStrategy struct{}

func (w *WalkRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Расчет пешего маршрута из %s в %s", start, end)
}

type PublicTransportRouteStrategy struct{}

func (p *PublicTransportRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Расчет маршрута на общественном транспорте из %s в %s", start, end)
}

type BikeRouteStrategy struct{}

func (b *BikeRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Расчет маршрута на велосипеде из %s в %s", start, end)
}

type Navigator struct {
	strategy RouteStrategy
}

func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

func (n *Navigator) CalculateRoute(start, end string) string {
	if n.strategy == nil {
		return "Стратегия маршрутизации не установлена"
	}
	return n.strategy.CalculateRoute(start, end)
}

func main() {
	navigator := &Navigator{}

	navigator.SetStrategy(&CarRouteStrategy{})
	fmt.Println(navigator.CalculateRoute("Красная площадь", "Тропаревский парк"))

	navigator.SetStrategy(&WalkRouteStrategy{})
	fmt.Println(navigator.CalculateRoute("Красная площадь", "Тропаревский парк"))

	navigator.SetStrategy(&PublicTransportRouteStrategy{})
	fmt.Println(navigator.CalculateRoute("Красная площадь", "Тропаревский парк"))

	navigator.SetStrategy(&BikeRouteStrategy{})
	fmt.Println(navigator.CalculateRoute("Красная площадь", "Тропаревский парк"))
}
