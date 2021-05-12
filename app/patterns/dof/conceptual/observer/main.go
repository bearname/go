package main

func main() {
	shirtItem := newItem("Nike Shirt")
	customer1 := &Customer{id: "abc@gmail.com"}
	customer2 := &Customer{id: "abc+1@gmail.com"}
	shirtItem.register(customer1)
	shirtItem.register(customer2)
	shirtItem.updateAvailability()
}
