package patterns

import "fmt"

// Builder provides a builder interface.
type ComputerBuilder interface {
	buildCPU(x string)
	buildMemory(x int)
	buildStorage(x int)
}

//Product
type Computer struct {
	cpu     string
	memory  int
	storage int
}

func (computer *Computer) Show() {
	fmt.Println("Computer created. It's configuration:")
	fmt.Println("CPU: ", computer.cpu)
	fmt.Println("Memory: ", computer.memory, "GB")
	fmt.Println("Storage: ", computer.storage, "GB")
}

//manager
type Director struct {
	builder ComputerBuilder
}

// Construct tells the builder what to do and in what order.
func (d *Director) Construct(cpu string, memory int, storage int) {
	d.builder.buildCPU(cpu)
	d.builder.buildMemory(memory)
	d.builder.buildStorage(storage)

}

// ConcreteBuilder implements Builder interface.
type ConcreteBuilder struct {
	computer *Computer
}

func (b *ConcreteBuilder) buildCPU(cpu string) {
	b.computer.cpu = cpu
}

func (b *ConcreteBuilder) buildMemory(memory int) {
	b.computer.memory = memory
}

func (b *ConcreteBuilder) buildStorage(storage int) {
	b.computer.storage = storage
}

func main_builder() {
	computer := new(Computer)
	cpu := "Intel i9 14900K"
	memory := 512 //GB
	storage := 16 //GB
	director := Director{&ConcreteBuilder{computer}}
	director.Construct(cpu, memory, storage)
	computer.Show()

}
