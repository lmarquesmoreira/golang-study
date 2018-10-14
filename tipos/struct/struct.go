package tipos

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

func (p Product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}


func ExecuteDemo(){

	product := Product{"Dell Notebook", 1989.90, 0.10}

	fmt.Println(product, product.priceWithDiscount())

	product2 := Product{"Vivo Notebook", 1999.90, 0.10}

	fmt.Println(product2, product2.priceWithDiscount())
}
