package Solution_1

import (
	"reflect"
	"testing"
)

func Test_LRUCache(t *testing.T) {
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
				functions: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				arguments: [][]interface{}{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			want: []interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]interface{}, len(tt.args.functions))

			var lruCache LRUCache
			for i, function := range tt.args.functions {
				switch function {
				case "LRUCache":
					lruCache = Constructor(tt.args.arguments[i][0].(int))
					got[i] = nil
				case "put":
					lruCache.Put(tt.args.arguments[i][0].(int), tt.args.arguments[i][1].(int))
					got[i] = nil
				case "get":
					value := lruCache.Get(tt.args.arguments[i][0].(int))
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
