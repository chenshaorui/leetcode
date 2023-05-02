package Extension_4

import "testing"

func Test_findTheLastNumberEqualToTarget(t *testing.T) {
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLastNumberEqualToTarget(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTheLastNumberEqualToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
