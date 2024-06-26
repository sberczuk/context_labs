package pkg

//Functions for reporting Pulled out into functions for ease of testing/debugging
// NOTE: I skipped this test

// total Price is Price * quanity for each product
func TotalPrice(products []*Product) float64 {
	var price float64
	for _, p := range products {
		price = price + p.GetPrice()*float64(p.GetQuantity())
	}

	return price
}

// Average Price is Price * Quantity/total Quantity
func AveragePrice(products []*Product) float64 {
	var price float64
	var count float64
	for _, p := range products {
		q := float64(p.GetQuantity())
		price = price + p.GetPrice()*q
		count = count + q
	}

	return price / count
}
