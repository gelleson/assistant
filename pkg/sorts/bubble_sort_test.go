package sorts

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func genShuffle(n int) ([]int, []int) {
	slc := make([]int, 0)
	orderedSlice := make([]int, 0)

	for i := 0; i < n; i++ {
		slc = append(slc, i)
		orderedSlice = append(orderedSlice, i)
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(slc), func(i, j int) {
		slc[i], slc[j] = slc[j], slc[i]
	})

	return orderedSlice, slc
}

func TestBubbleSort_Sort(t *testing.T) {

	ordered, shuffled := genShuffle(100)
	type args struct {
		un []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort",
			args: args{
				un: shuffled,
			},
			want: ordered,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BubbleSort{}
			if got := s.Sort(tt.args.un); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort_Sort(b *testing.B) {
	_, shuffled := genShuffle(100)
	b.ResetTimer()
	bubbleSorter := BubbleSort{}
	for i := 0; i < b.N; i++ {
		bubbleSorter.Sort(shuffled)
	}
}
