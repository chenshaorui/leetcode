package Extension_1

import "testing"

func Test_findTheFirstNumberEqualToTarget(t *testing.T) {
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
				nums:   []int{1, 2, 3, 3, 4, 5, 6},
				target: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheFirstNumberEqualToTarget(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTheFirstNumberEqualToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
