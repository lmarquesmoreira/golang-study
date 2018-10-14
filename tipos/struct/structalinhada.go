package tipos

import "fmt"

type Item struct {
	productId int
	quantity  int
	price     float64
}

type Order struct {
	userId int
	items  []Item
}

func (o Order) totalPrice() (total float64) {
	total = 0.0
	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}
	return
}

func ExecuteDemoAlinhada() {
	order := Order{
		userId: 1, items: []Item{
			Item{productId: 1, quantity: 2, price: 12.10},
			Item{2, 1, 23.49},
			Item{11, 100, 3.13},
		},
	}

	fmt.Printf("Order price is U$ %.2f\n", order.totalPrice())
}
