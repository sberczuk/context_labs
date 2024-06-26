package pkg

type Rule struct {
	Conditions []Condition `json:"conditions,omitempty"`
	Score      int         `json:"score,omitempty"`
}

func NewRule(conditions []Condition, score int) *Rule {
	return &Rule{
		conditions,
		score,
	}
}

func (r *Rule) AddCondition(condition Condition) {
	r.Conditions = append(r.Conditions, condition)
}

type RuleResult struct {
	NumMatches      int
	TotalConditions int
	Score           float64
}

func (r *Rule) Evaluate(product *Product) RuleResult {
	var matchedConditions int
	for _, c := range r.Conditions {
		satisfies := c.Apply(product)
		if satisfies {
			matchedConditions++
		}
	}

	// return match status and percent
	totalConditions := len(r.Conditions)

	score := float64(r.Score) * float64(matchedConditions) / float64(totalConditions)

	// I could also just count passed conditions, but this seems more intuitive
	return RuleResult{
		NumMatches:      matchedConditions,
		TotalConditions: totalConditions,
		Score:           score,
	}
}

type RuleSet struct {
	Rules []Rule
}

func NewRuleSet(rules []Rule) *RuleSet {
	return &RuleSet{
		Rules: rules,
	}
}

//func (s *RuleSet) Score(product Product) int {
//	var score int
//	for _, rule := range s.Rules {
//		_, i := rule.Score(product)
//		score = score + i
//	}
//	return score
//}

type ColorAndPriceRule struct {
}

//func (r *ColorAndPriceRule) Matches(p Product) bool {
//	Color := p.GetAttribute("COLOR")
//	Price := p.GetAttribute("PRICE")
//	// add a check to confirm the type?
//	return Color.Equals("BLUE") && Price.LessThan(100.22)
//}

//func (r *ColorAndPriceRule) Score(p Product) (bool, int) {
//	if r.Matches(p) {
//		return true, 100
//	}
//	return false, 0
//}
