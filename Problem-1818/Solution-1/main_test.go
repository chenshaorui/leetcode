package Solution_1

import "testing"

func Test_minAbsoluteSumDiff(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test",
			args: args{
				nums1: []int{1, 7, 5},
				nums2: []int{2, 3, 5},
			},
			want: 3,
		},
		{
			name: "Test",
			args: args{
				nums1: []int{2, 4, 6, 8, 10},
				nums2: []int{2, 4, 6, 8, 10},
			},
			want: 0,
		},
		{
			name: "Test",
			args: args{
				nums1: []int{1, 10, 4, 4, 2, 7},
				nums2: []int{9, 3, 5, 1, 7, 4},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsoluteSumDiff(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minAbsoluteSumDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
