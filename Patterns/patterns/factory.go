package patterns

import (
	"fmt"
)

// Product интерфейс
type Product interface{}

// ConcreteProductA структура
type ConcreteProductA struct{}

// ConcreteProductB структура
type ConcreteProductB struct{}

// Creator интерфейс
type Creator interface {
	FactoryMethod() Product
}

// ConcreteCreatorA структура
type ConcreteCreatorA struct{}

func (c ConcreteCreatorA) FactoryMethod() Product {
	return ConcreteProductA{}
}

// ConcreteCreatorB структура
type ConcreteCreatorB struct{}

func (c ConcreteCreatorB) FactoryMethod() Product {
	return ConcreteProductB{}
}

func main_factory() {
	// массив создателей
	creators := []Creator{ConcreteCreatorA{}, ConcreteCreatorB{}}
	// итерация по создателям и создание продуктов
	for _, creator := range creators {
		product := creator.FactoryMethod()
		fmt.Printf("Created {%T}\n", product)
	}
}
