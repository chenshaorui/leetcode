package Extension_2

import "testing"

func Test_findTheFirstNumberGreaterThanOrEqualToTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test",
			args: args{
				nums:   []int{1, 2, 4, 5, 6},
				target: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheFirstNumberGreaterThanOrEqualToTarget(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTheFirstNumberGreaterThanOrEqualToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
