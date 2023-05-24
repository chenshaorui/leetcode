package Solution_1

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		intersectVal int
		headA        *ListNode
		headB        *ListNode
		skipA        int
		skipB        int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test",
			args: args{
				intersectVal: 8,
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				headB: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  5,
										Next: nil,
									},
								},
							},
						},
					},
				},
				skipA: 2,
				skipB: 3,
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Test",
			args: args{
				intersectVal: 2,
				headA: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
				headB: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				skipA: 3,
				skipB: 1,
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		{
			name: "Test",
			args: args{
				intersectVal: 0,
				headA: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				headB: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
				skipA: 3,
				skipB: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.intersectVal > 0 {
				currA := tt.args.headA
				for i := 0; i < tt.args.skipA && currA != nil; i++ {
					currA = currA.Next
				}

				currB := tt.args.headB
				for i := 1; i < tt.args.skipB && currB != nil; i++ {
					currB = currB.Next
				}

				currB.Next = currA
			}

			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
