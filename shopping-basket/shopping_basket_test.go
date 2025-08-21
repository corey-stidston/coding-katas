package main

import (
	"fmt"
	"testing"
)

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
	tests := []struct {
		testName      string
		items         []Item
		expectedPrice float32
	}{
		{
			testName: "Basket value less than $100",
			items: []Item{
				{
					Name:  "orange",
					Price: 1.50,
				},
				{
					Name:  "apple",
					Price: 2.25,
				},
			},
			expectedPrice: 3.75,
		},
		{
			testName: "Basket value greater than $100, less than $200",
			items: []Item{
				{
					Name:  "truffles",
					Price: 75.0,
				},
				{
					Name:  "olive oil",
					Price: 50,
				},
			},
			expectedPrice: 118.75,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			basket := ShoppingBasket{
				Basket: tt.items,
			}

			totalPrice := basket.GetTotalPrice()

			if totalPrice != tt.expectedPrice {
				t.Errorf("Expected a total price of %f but got %f", tt.expectedPrice, totalPrice)
			}
		})
	}
}
