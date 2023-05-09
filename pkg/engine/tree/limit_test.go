package tree_test

import (
	"kwil/pkg/engine/tree"
	"testing"
)

func TestLimit_ToSQL(t *testing.T) {
	type fields struct {
		Expression       tree.Expression
		Offset           tree.Expression
		SecondExpression tree.Expression
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		wantPanic bool
	}{
		{
			name: "valid limit",
			fields: fields{
				Expression: &tree.ExpressionBindParameter{Parameter: "$a"},
			},
			want: ` LIMIT $a`,
		},
		{
			name: "valid limit with offset",
			fields: fields{
				Expression: &tree.ExpressionBindParameter{Parameter: "$a"},
				Offset:     &tree.ExpressionBindParameter{Parameter: "$b"},
			},
			want: ` LIMIT $a OFFSET $b`,
		},
		{
			name: "valid limit with second expression",
			fields: fields{
				Expression: &tree.ExpressionBindParameter{Parameter: "$a"},
				SecondExpression: &tree.ExpressionBinaryComparison{
					Left:     &tree.ExpressionColumn{Column: "foo"},
					Operator: tree.ComparisonOperatorEqual,
					Right:    &tree.ExpressionBindParameter{Parameter: "$b"},
				},
			},
			want: ` LIMIT $a, "foo" = $b`,
		},
		{
			name:      "no expression",
			fields:    fields{},
			wantPanic: true,
		},
		{
			name: "offset and second expression",
			fields: fields{
				Expression: &tree.ExpressionBindParameter{Parameter: "$a"},
				Offset:     &tree.ExpressionBindParameter{Parameter: "$b"},
				SecondExpression: &tree.ExpressionBinaryComparison{
					Left:     &tree.ExpressionColumn{Column: "foo"},
					Operator: tree.ComparisonOperatorEqual,
					Right:    &tree.ExpressionBindParameter{Parameter: "$c"},
				},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &tree.Limit{
				Expression:       tt.fields.Expression,
				Offset:           tt.fields.Offset,
				SecondExpression: tt.fields.SecondExpression,
			}

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Limit.ToSQL() should have panicked")
					}
				}()
			}

			got := l.ToSQL()
			if tt.wantPanic {
				return
			}

			if !compareIgnoringWhitespace(got, tt.want) {
				t.Errorf("Limit.ToSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}