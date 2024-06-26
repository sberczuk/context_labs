package pkg

type Score struct {
	p               Product
	r               []Rule
	numConditions   int
	matchConditions int
	score           int
}

func NewScore(p Product, r []Rule) *Score {
	return &Score{p: p, r: r}
}

func (s *Score) sumOfRulesScore() float64 {
	return float64(s.matchConditions) / float64(len(s.r)) * float64(s.score)
}
