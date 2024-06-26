package pkg

import "testing"

func TestFilterProduct_Apply(t *testing.T) {
	type fields struct {
		MaxScoreThreshold float64
	}
	type args struct {
		p *Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Pass",
			fields: fields{
				MaxScoreThreshold: 0.5,
			},
			args: args{
				p: &Product{
					maxPossibleScore: 200,
					productScore:     100,
				},
			},
			want: true,
		},

		{
			name: "Fail",
			fields: fields{
				MaxScoreThreshold: 0.5,
			},
			args: args{
				p: &Product{
					maxPossibleScore: 500,
					productScore:     100,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FilterProduct{
				MaxScoreThreshold: tt.fields.MaxScoreThreshold,
			}
			if got := f.Apply(tt.args.p); got != tt.want {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
