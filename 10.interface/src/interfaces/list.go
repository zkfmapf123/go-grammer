package interfaces

type Printer interface {
	Print()
}

type List []Printer

func (l List) Print() {

	for _, it := range l {
		it.Print()
	}
}

func (l List) Discount(ratio float64) {

	type discountner interface {
		Discount(float64)
	}

	for _, it := range l {
		if it, ok := it.(discountner); ok {
			it.Discount(ratio)
		}
	}
}
