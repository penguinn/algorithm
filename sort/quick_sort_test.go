package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{data: []int{1000, 2, 31, 34, 5, 9, 7, 4, 6, 89, 90, 99, 99, 99, 99, 99}},
			want: []int{2, 4, 5, 6, 7, 9, 31, 34, 89, 90, 99, 99, 99, 99, 99, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
