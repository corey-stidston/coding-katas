package main

import "testing"

func TestAddItemsToShoppingBasket(t *testing.T) {
	basket := ShoppingBasket {}
	basket.AddItem("orange")

	if(len(basket.Basket) != 1 || basket.Basket[0] != "orange") {
		t.Errorf("Expected basket to contain exactly 1 item that is an orange")
	}
}

func TestGetQuantity(t *testing.T) {
	basket := ShoppingBasket {
		Basket: []string{"orange", "orange"},
	}

	expected := 2
	numberOfOranges := basket.GetItemQuantity("orange")

	if (numberOfOranges != expected) {
		t.Errorf("Expected an item quantity of %d oranges", expected)
	}
}
