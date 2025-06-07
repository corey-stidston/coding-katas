/*
Items in a shopping basket have a unit price and quantity. Write code that will allow you to:

- find out the quantity of a particular item in the basket
- calculate the total price of the whole basket, including any applicable discount

Normally the total price is the sum of unit price * quantity for all the items.
If you buy in bulk you get a discount:

If total basket value > $100, apply a 5% discount
If total basket value > $200, apply a 10% discount

Example
Item A: price $10, quantity 5
Item B: price $25, quantity 2
Item C: price $9.99, quantity 6
This basket qualifies for a 5% discount and the total price is $151.94
*/

package main

type Item struct {
	Name  string
	Price float32
}

type ShoppingBasket struct {
	Basket []Item
}

func (sb *ShoppingBasket) AddItem(item Item) {
	sb.Basket = append(sb.Basket, item)
}

func (sb *ShoppingBasket) GetTotalPrice() float32 {
	totalPrice := float32(0)
	for _, item := range sb.Basket {
		totalPrice += item.Price
	}

	discount := getDiscount(totalPrice)

	return totalPrice - (totalPrice * discount)
}

func getDiscount(totalPrice float32) float32 {
	if totalPrice < 100 {
		return 0
	} else if totalPrice > 100 && totalPrice < 200 {
		return 0.05
	} else { // totalPrice > 200
		return 0.1
	}
}

func (sb ShoppingBasket) GetItemQuantity(itemName string) int {
	count := 0
	for _, item := range sb.Basket {
		if item.Name == itemName {
			count++
		}
	}
	return count
}

func main() {
	//
}
