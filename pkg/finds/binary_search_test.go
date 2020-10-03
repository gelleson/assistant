package finds

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		searchTerm int
		arr        []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "binary_search",
			args: args{
				searchTerm: 4,
				arr:        []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.searchTerm, tt.args.arr); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
