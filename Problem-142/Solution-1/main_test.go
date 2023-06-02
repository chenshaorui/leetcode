package Solution_1

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
		pos  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
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
			want: &ListNode{
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
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "Test",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				pos: 0,
			},
			want: nil,
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

			if tt.want != nil {
				curr := tt.want

				for curr.Next != nil {
					curr = curr.Next
				}

				curr.Next = tt.want
			}

			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
