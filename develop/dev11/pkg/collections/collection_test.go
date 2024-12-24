package collections

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		collection []any
		predicate  func(any) bool
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "filter",
			args: args{
				collection: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				predicate: func(i any) bool {
					return i.(int) > 5
				},
			},
			want: []any{6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.collection, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
