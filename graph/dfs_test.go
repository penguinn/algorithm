package graph

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	type args struct {
		graph       *Graph
		startIndex  int
		targetIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				graph:       &Graph{Nodes: map[int][]int{0: {1, 2}, 1: {3, 4}, 2: {5, 6, 7}, 3: {7}, 4: {}, 5: {}, 6: {}, 7: {}}},
				startIndex:  0,
				targetIndex: 7,
			},
			want: []int{0, 1, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DFS(tt.args.graph, tt.args.startIndex, tt.args.targetIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
