package Solution_1

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  -4,
								Next: nil,
							},
						},
					},
				},
				pos: 1,
			},
			want: true,
		},
		{
			name: "Test",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				pos: 0,
			},
			want: true,
		},
		{
			name: "Test",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				pos: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.pos >= 0 && tt.args.head != nil {
				cycle, curr := (*ListNode)(nil), tt.args.head

				for i := 0; curr.Next != nil; i++ {
					if i == tt.args.pos {
						cycle = curr
					}

					curr = curr.Next
				}

				curr.Next = cycle
			}

			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
