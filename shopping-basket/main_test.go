package main

import "testing"

func TestAddItemsToShoppingBasket(t *testing.T) {
	basket := Build()
	basket.AddItem("orange")

	if(len(basket.Basket) != 1 || basket.Basket[0] != "orange") {
		t.Errorf("Expected basket to contain exactly 1 item that is an orange")
	}
}
