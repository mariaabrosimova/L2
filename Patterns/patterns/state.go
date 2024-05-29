package patterns

import (
	"fmt"
)

// VendingMachineState интерфейс
type VendingMachineState interface {
	HandleRequest()
}

// ReadyState структура
type ReadyState struct{}

func (r ReadyState) HandleRequest() {
	fmt.Println("Ready state: Please select a product.")
}

// ProductSelectedState структура
type ProductSelectedState struct{}

func (p ProductSelectedState) HandleRequest() {
	fmt.Println("Product selected state: Processing payment.")
}

// PaymentPendingState структура
type PaymentPendingState struct{}

func (p PaymentPendingState) HandleRequest() {
	fmt.Println("Payment pending state: Dispensing product.")
}

// OutOfStockState структура
type OutOfStockState struct{}

func (o OutOfStockState) HandleRequest() {
	fmt.Println("Out of stock state: Product unavailable. Please select another product.")
}

func main_state() {
	var state VendingMachineState

	state = ReadyState{}
	state.HandleRequest()

	state = ProductSelectedState{}
	state.HandleRequest()

	state = PaymentPendingState{}
	state.HandleRequest()

	state = OutOfStockState{}
	state.HandleRequest()
}
