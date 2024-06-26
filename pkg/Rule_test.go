package pkg

import (
	"testing"
)

var blueColor = BaseAttribute[string]{
	Name:  "COLOR",
	Value: "BLUE",
}

func TestRule_Evaluate(t *testing.T) {
	type fields struct {
		Conditions []Condition
	}
	type args struct {
		product *Product
	}
	var tests = []struct {
		name           string
		fields         fields
		args           args
		wantNumMatches int
		wantScore      float64
	}{
		{
			name: "Product Matches all conditions",
			fields: fields{
				Conditions: []Condition{NewConcreteCondition[string]("COLOR", &EqualOperator[string]{}, "BLUE")},
			},
			args: args{
				product: &Product{
					attribMap: map[AttribName]Attribute{Color: blueColor},
				},
			},
			wantNumMatches: 1,
			wantScore:      100,
		},
		{
			name: "Product Matches No conditions",
			fields: fields{
				Conditions: []Condition{NewConcreteCondition[string](Color, &EqualOperator[string]{}, "RED")},
			},
			args: args{
				product: &Product{
					attribMap: map[AttribName]Attribute{Color: blueColor},
				},
			},
			wantNumMatches: 0,
			wantScore:      0,
		},

		{
			name: "Product Matches 1 of 4 conditions",
			fields: fields{
				Conditions: []Condition{
					NewConcreteCondition[string](Color, &EqualOperator[string]{}, "RED"),
					NewConcreteCondition[string](Color, &EqualOperator[string]{}, "BLUE"),
					NewConcreteCondition[string](Color, &EqualOperator[string]{}, "ORANGE"),
					NewConcreteCondition[string](Color, &EqualOperator[string]{}, "YELLOW"),
				},
			},
			args: args{
				product: &Product{
					attribMap: map[AttribName]Attribute{Color: blueColor},
				},
			},
			wantNumMatches: 1,
			wantScore:      25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rule{
				Conditions: tt.fields.Conditions,
				Score:      100,
			}
			got := r.Evaluate(tt.args.product)
			if got.Score != tt.wantScore {
				t.Errorf("Evaluate() got = %v, want %v", got.Score, tt.wantScore)
			}
			if got.NumMatches != tt.wantNumMatches {
				t.Errorf("Evaluate() got1 = %v, want %v", got.NumMatches, tt.wantNumMatches)
			}
		})
	}
}
