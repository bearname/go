package main

type tomatoTopping struct {
	pizza pizza
}

func (t tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}
