package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

Паттерн применяется, когда нужно спрятать сложный функционал в один интерфейс, с которым пользователь будет взаимодействовать.

Плюсы:
	1 Простой интерфейс для взаимодействия
	2 Инкапсуляция

Минусы:
	1 Дополнительный слой абстракции
	2 Слишком большой функционал в единой точке

*/

type CPU struct{}

func (cpu *CPU) Power() {
	fmt.Println("CPU Power ON")
}

type GPU struct{}

func (gpu *GPU) Power() {
	fmt.Println("GPU Power ON")
}

type RAM struct{}

func (ram *RAM) Power() {
	fmt.Println("RAM Power ON")
}

type SSD struct{}

func (sd *SSD) Power() {
	fmt.Println("SSD Power ON")
}

type Motherboard struct {
	cpu CPU
	gpu GPU
	ram RAM
	ssd SSD
}

func NewMotherboard() *Motherboard {
	return &Motherboard{
		cpu: CPU{},
		gpu: GPU{},
		ram: RAM{},
		ssd: SSD{},
	}
}

func (mb *Motherboard) Power() {
	mb.cpu.Power()
	mb.gpu.Power()
	mb.ram.Power()
	mb.ssd.Power()
}

func main() {
	pc := NewMotherboard()
	pc.Power()
}
