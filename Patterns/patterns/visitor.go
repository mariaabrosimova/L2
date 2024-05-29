package patterns

import "fmt"

type Visitor interface {
	VisitForInch(w Inch)
	VisitForFoot(w Foot)
}

// элемент
type Length interface {
	accept(v Visitor)
}

// конкретные элементы
type Inch struct {
	length float32
}

// метод accept, который вызывает необходимый метод посещения
func (w *Inch) accept(v Visitor) {
	v.VisitForInch(*w)
}

type Foot struct {
	length float32
}

func (w *Foot) accept(v Visitor) {
	v.VisitForFoot(*w)
}

// конкретные посетители
type ConvertVisitor struct {
}

func (cv *ConvertVisitor) VisitForInch(w Inch) {
	fmt.Printf("Length = %.4f m\n", w.length*0.0254)
}

func (cv *ConvertVisitor) VisitForFoot(w Foot) {
	fmt.Printf("Length = %.4f m\n", w.length*0.3048)
}

func main_visitor() {
	Inch := &Inch{length: 50}
	Foot := &Foot{length: 50}

	lenghts_list := []Length{Inch, Foot}
	fmt.Printf("Длина \nдюйм: %.4f \nфут:%.4f\n", Inch.length, Foot.length)
	fmt.Println("Конвертация в метры")
	ucv := &ConvertVisitor{}
	//вызываем метод accept который у каждой структуры вызовет метод для своего типа
	for _, w := range lenghts_list {
		w.accept(ucv)
	}

}
