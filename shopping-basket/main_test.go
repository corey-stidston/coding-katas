package main

import "testing"

func TestAddItemsToShoppingBasket(t *testing.T) {
	basket := ShoppingBasket{}
	basket.AddItem(Item{
		Name:  "orange",
		Price: 1.50,
	})

	if (len(basket.Basket) != 1 || basket.Basket[0] != Item{Name: "orange", Price: 1.50}) {
		t.Errorf("Expected basket to contain exactly 1 item that is an orange")
	}
}

func TestGetQuantity(t *testing.T) {
	basket := ShoppingBasket{
		Basket: []Item{{
			Name:  "orange",
			Price: 1.50,
		}, {
			Name:  "orange",
			Price: 1.50,
		}},
	}

	expected := 2
	numberOfOranges := basket.GetItemQuantity("orange")

	if numberOfOranges != expected {
		t.Errorf("Expected an item quantity of %d oranges", expected)
	}
}
