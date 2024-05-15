package main

import "fmt"

/*
	Строитель позволяет создавать сложные объекты пошагово.
	Это позволяет использовать один и тот же код для получения разных представлений объектов

Плюсы:
	1 Позволяет создавать продукты пошагово.
 	2 Позволяет использовать один и тот же код для создания различных продуктов.
	3 Изолирует сложный код сборки продукта от его основной бизнес-логики.
Минусы
	1 Усложняет код программы из-за введения дополнительных классов.
	2 Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

type CPU struct {
	Model        string
	Cores        int
	Threads      int
	BaseClock    float64
	BoostClock   float64
	Architecture string
}

type CPUBuilder interface {
	SetModel(model string) CPUBuilder
	SetCores(cores int) CPUBuilder
	SetThreads(threads int) CPUBuilder
	SetBaseClock(clock float64) CPUBuilder
	SetBoostClock(clock float64) CPUBuilder
	SetArchitecture(arch string) CPUBuilder
	Build() CPU
}

type ConcreteCPUBuilder struct {
	cpu CPU
}

func NewCPUBuilder() *ConcreteCPUBuilder {
	return &ConcreteCPUBuilder{}
}

func (b *ConcreteCPUBuilder) SetModel(model string) CPUBuilder {
	b.cpu.Model = model
	return b
}

func (b *ConcreteCPUBuilder) SetCores(cores int) CPUBuilder {
	b.cpu.Cores = cores
	return b
}

func (b *ConcreteCPUBuilder) SetThreads(threads int) CPUBuilder {
	b.cpu.Threads = threads
	return b
}

func (b *ConcreteCPUBuilder) SetBaseClock(clock float64) CPUBuilder {
	b.cpu.BaseClock = clock
	return b
}

func (b *ConcreteCPUBuilder) SetBoostClock(clock float64) CPUBuilder {
	b.cpu.BoostClock = clock
	return b
}

func (b *ConcreteCPUBuilder) SetArchitecture(arch string) CPUBuilder {
	b.cpu.Architecture = arch
	return b
}

func (b *ConcreteCPUBuilder) Build() CPU {
	return b.cpu
}

type Director struct {
	builder CPUBuilder
}

func NewDirector(builder CPUBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) ConstructHighEndCPU() CPU {
	return d.builder.
		SetModel("Ryzen 9 5900X").
		SetCores(16).
		SetThreads(32).
		SetBaseClock(3.4).
		SetBoostClock(4.9).
		SetArchitecture("Zen 3").
		Build()
}

func (d *Director) ConstructMidRangeCPU() CPU {
	return d.builder.
		SetModel("Ryzen 5 5600X").
		SetCores(6).
		SetThreads(12).
		SetBaseClock(3.7).
		SetBoostClock(4.6).
		SetArchitecture("Zen 3").
		Build()
}

func main() {
	builder := NewCPUBuilder()
	director := NewDirector(builder)

	highEndCPU := director.ConstructHighEndCPU()
	fmt.Printf("High-end CPU: %+v\n", highEndCPU)

	midRangeCPU := director.ConstructMidRangeCPU()
	fmt.Printf("Mid-range CPU: %+v\n", midRangeCPU)
}
