package Solution_1

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Test",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   4,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   8,
							Left:  nil,
							Right: nil,
						},
					},
				},
				p: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "Test",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   4,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   8,
							Left:  nil,
							Right: nil,
						},
					},
				},
				p: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
