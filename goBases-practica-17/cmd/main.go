package main

import "fmt"

const (
	sm = "Small"
	md = "Medium"
	lg = "Large"
)

type Small struct {
	Price float64
}

func (s Small) price() float64 {
	return s.Price
}

type Medium struct {
	Price float64
}

func (m Medium) price() float64 {
	return m.Price * 1.03
}

type Large struct {
	Price float64
}

func (l Large) price() float64 {
	return l.Price*1.06 + 2500
}

type Product interface {
	price() float64
}

func factory(productType string, price float64) Product {
	switch productType {
	case sm:
		return Small{Price: price}

	case md:
		return Medium{Price: price}

	case lg:
		return Large{Price: price}
	}

	return nil
}

func main() {

	smallP := factory(sm, 120.0)
	fmt.Printf("\nPrice: $%.2f\n\n", smallP.price())

	mediumP := factory(md, 120.0)
	fmt.Printf("Price: $%.2f\n\n", mediumP.price())

	largeP := factory(lg, 120.0)
	fmt.Printf("Price: $%.2f\n\n", largeP.price())

}
