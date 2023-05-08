package Solution_1

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name: "Test",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "Test",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
