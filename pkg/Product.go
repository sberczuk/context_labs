package pkg

type Product struct {
	attribMap map[AttribName]Attribute `json:"attrib_map,omitempty"`

	// These fields are populated when the product is processed.
	// we could also wrap products in another struct as we filter.
	// This is simpler for the moment
	numConditionsApplied int
	numConditionsMatched int
	productScore         float64
	maxPossibleScore     float64
}

func NewProduct(a []Attribute) *Product {
	am := make(map[AttribName]Attribute)
	for _, a := range a {
		am[a.GetName()] = a
	}
	return &Product{
		attribMap: am,
	}
}

func (p *Product) GetMatchFraction() float64 {
	return float64(p.productScore) / float64(p.maxPossibleScore)
}

// Score Scores the Product Based on the Rules
func (p *Product) Score(rules []Rule) {
	for _, r := range rules {
		ruleResult := r.Evaluate(p)
		p.numConditionsApplied = ruleResult.TotalConditions
		p.numConditionsMatched = ruleResult.NumMatches
		p.productScore += ruleResult.Score
		p.maxPossibleScore = p.maxPossibleScore + float64(r.Score)
	}
}

func (p *Product) AddAttribute(a Attribute) {
	//this isn't type safe as is, but we could add a function to convert in an error safe way
	// there is also a go-enum tool that will bgenerate the code, but I'm ommitting the for brevity
	// given that the Attributes are parameterized,  this seems reasonably safe for a prototype
	p.attribMap[AttribName(a.GetName())] = a
}

func (p *Product) GetAttribute(name AttribName) Attribute {
	return p.attribMap[name]
}

func (p *Product) GetPrice() float64 {
	return p.GetAttribute(Price).GetValue().(float64)
}

func (p *Product) GetQuantity() int {
	return p.GetAttribute(Quantity).GetValue().(int)

}
