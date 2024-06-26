package main

import (
	"context_labs/pkg"
	"fmt"
)

func getProducts() ([]*pkg.Product, error) {
	p1 := pkg.NewProduct([]pkg.Attribute{
		pkg.BaseAttribute[float64]{Name: pkg.Price, Value: 100.23},
		pkg.BaseAttribute[int]{Name: pkg.Quantity, Value: 70},
		pkg.BaseAttribute[string]{Name: pkg.Color, Value: "BLUE"},
	})
	p2 := pkg.NewProduct([]pkg.Attribute{
		pkg.BaseAttribute[float64]{Name: pkg.Price, Value: 100.23},
		pkg.BaseAttribute[int]{Name: pkg.Quantity, Value: 70},
		pkg.BaseAttribute[string]{Name: pkg.Color, Value: "BLUE"},
	})

	p3 := pkg.NewProduct([]pkg.Attribute{
		pkg.BaseAttribute[float64]{Name: pkg.Price, Value: 100.23},
		pkg.BaseAttribute[int]{Name: pkg.Quantity, Value: 70},
		pkg.BaseAttribute[string]{Name: pkg.Color, Value: "BLUE"},
	})

	products := []*pkg.Product{p1, p2, p3}
	return products, nil
}

func getRules() []pkg.Rule {
	r1 := pkg.Rule{
		Conditions: []pkg.Condition{
			&pkg.ConcreteCondition[float64]{pkg.Price, &pkg.LessThanOperator[float64]{}, 100.56},
			&pkg.ConcreteCondition[string]{pkg.Color, &pkg.EqualOperator[string]{}, "BLUE"},
		},
		Score: 100,
	}

	r2 := pkg.Rule{
		Conditions: []pkg.Condition{
			&pkg.ConcreteCondition[int]{pkg.Quantity, &pkg.LessThanOperator[int]{}, 200},
			&pkg.ConcreteCondition[string]{pkg.Color, &pkg.EqualOperator[string]{}, "BLUE"},
		},
		Score: 100,
	}

	rules := []pkg.Rule{r1, r2}
	return rules
}

func main() {

	products, _ := getProducts()
	rules := getRules()
	// I could use NewRule for this, but he structs seem more evocative

	// apply rules
	for _, product := range products {
		product.Score(rules)
	}

	//filter
	filter := pkg.FilterProduct{MaxScoreThreshold: 0.5}
	var filtered []*pkg.Product

	for _, product := range products {
		if filter.Apply(product) {
			filtered = append(filtered, product)
		}
	}
	//  report

	fmt.Printf("Total Price %f Avg Price %f", pkg.TotalPrice(filtered), pkg.AveragePrice(filtered))

}
