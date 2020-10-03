package sorts

import (
	"reflect"
	"testing"
)

func TestCocktailSort_Sort(t *testing.T) {

	ordered, unordered := genShuffle(100)

	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "cocktail sort",
			args: args{
				arr: unordered,
			},
			want: ordered,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CocktailSort{}
			if got := c.Sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
