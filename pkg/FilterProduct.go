package pkg

// functions to do filtering.
// pulled out in a new function for ease of testing

type FilterProduct struct {
	MaxScoreThreshold float64
}

func (f *FilterProduct) Apply(p *Product) bool {
	matchFraction := p.GetMatchFraction()
	return matchFraction >= f.MaxScoreThreshold
}
