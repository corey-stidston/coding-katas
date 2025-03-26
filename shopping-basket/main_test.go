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
	orange := Item{
		Name:  "orange",
		Price: 1.50,
	}

	basket := ShoppingBasket{
		Basket: []Item{orange, orange},
	}

	expected := 2
	numberOfOranges := basket.GetItemQuantity("orange")

	if numberOfOranges != expected {
		t.Errorf("Expected an item quantity of %d oranges but got %d", expected, numberOfOranges)
	}
}

func TestGetTotalPrice(t *testing.T) {
	orange := Item{
		Name:  "orange",
		Price: 1.50,
	}
	apple := Item{
		Name: "apple",
		Price: 2.25,
	}

	basket := ShoppingBasket{
		Basket: []Item{orange, apple},
	}

	expected := float32(3.75)
	totalPrice := basket.GetTotalPrice()

	if totalPrice != expected {
		t.Errorf("Expected a total price of %f but got %f", expected, totalPrice)
	}
}
