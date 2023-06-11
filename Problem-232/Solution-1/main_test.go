package Solution_1

import (
	"reflect"
	"testing"
)

func Test_MyQueue(t *testing.T) {
	type args struct {
		functions []string
		arguments [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Test",
			args: args{
				functions: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
				arguments: [][]interface{}{{}, {1}, {2}, {}, {}, {}},
			},
			want: []interface{}{nil, nil, nil, 1, 1, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]interface{}, len(tt.args.functions))

			var myQueue MyQueue
			for i, function := range tt.args.functions {
				switch function {
				case "MyQueue":
					myQueue = Constructor()
					got[i] = nil
				case "push":
					myQueue.Push(tt.args.arguments[i][0].(int))
					got[i] = nil
				case "peek":
					value := myQueue.Peek()
					got[i] = value
				case "pop":
					value := myQueue.Pop()
					got[i] = value
				case "empty":
					value := myQueue.Empty()
					got[i] = value
				default:
					panic("invalid function")
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
