package Solution_1

import (
	"reflect"
	"sort"
	"testing"
)

type Results [][]int

var _ sort.Interface = (*Results)(nil)

func (results Results) Len() int {
	return len(results)
}

func (results Results) Swap(i, j int) {
	results[i], results[j] = results[j], results[i]
}

func (results Results) Less(i, j int) bool {
	len1, len2 := len(results[i]), len(results[j])

	if len1 < len2 {
		return true
	} else if len1 > len2 {
		return false
	} else {
		for k := 0; k < len1; k++ {
			if results[i][k] < results[j][k] {
				return true
			} else if results[i][k] > results[j][k] {
				return false
			}
		}
	}

	return false
}

func sortResults(results [][]int) [][]int {
	for _, result := range results {
		sort.Ints(result)
	}

	sort.Sort(Results(results))

	return results
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Test",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "Test",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(sortResults(got), sortResults(tt.want)) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
