package pkg

import "testing"

var blueColor2 = BaseAttribute[string]{
	Name:  Color,
	Value: "BLUE",
}

var price1 = BaseAttribute[float64]{
	Name:  Price,
	Value: 100,
}

var price2 = BaseAttribute[float64]{
	Name:  Price,
	Value: 200,
}

var priceCondition = ConcreteCondition[float64]{
	AttributeName: Price,
	Operator:      &EqualOperator[float64]{},
	Value:         100,
}

func TestProduct_Score(t *testing.T) {
	type fields struct {
		attribMap            map[AttribName]Attribute
		Price                float64
		NumConditionsApplied int
		NumConditionsMatched int
		ProductScore         int
	}
	type args struct {
		rules []Rule
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantScore    float64
		wantMaxScore float64
	}{
		{
			name: "Matches All",
			fields: fields{
				attribMap: map[AttribName]Attribute{Color: blueColor2, Price: price1},
			},
			args: args{
				rules: []Rule{{
					Conditions: []Condition{&priceCondition},
					Score:      100,
				}},
			},
			wantScore:    100,
			wantMaxScore: 100,
		},
		{
			name: "Matches None",
			fields: fields{
				attribMap: map[AttribName]Attribute{Color: blueColor2, Price: price2},
			},
			args: args{
				rules: []Rule{{
					Conditions: []Condition{&priceCondition},
					Score:      100,
				}},
			},
			wantScore:    0,
			wantMaxScore: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				attribMap: tt.fields.attribMap,
			}

			p.Score(tt.args.rules)
			if p.productScore != tt.wantScore {
				t.Errorf("Score() got1 = %v, want %v", p.productScore, tt.wantScore)
			}
			if p.maxPossibleScore != tt.wantMaxScore {
				t.Errorf("Score() got1 = %v, want %v", p.maxPossibleScore, tt.wantMaxScore)
			}
		})
	}
}
