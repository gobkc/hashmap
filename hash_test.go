package hashmap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHash_Get(t *testing.T) {
	type fields struct {
		bucketNum int64
		buckets   []*Bucket
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name: "example 1",
			fields: fields{
				bucketNum: 0,
				buckets:   nil,
			},
			args: args{
				key:   "test key 1",
				value: "test value 1",
			},
			want: &Node{
				K:       "test key 1",
				V:       "test value 1",
				deleted: false,
			},
		},
		{
			name: "example 2",
			fields: fields{
				bucketNum: 0,
				buckets:   nil,
			},
			args: args{
				key:   "test key 2",
				value: "test value 2",
			},
			want: &Node{
				K:       "test key 2",
				V:       "test value 2",
				deleted: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hash{
				bucketNum: tt.fields.bucketNum,
				buckets:   tt.fields.buckets,
			}
			h.Set(tt.args.key, tt.args.value)
			if got := h.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash_Set(t *testing.T) {
	h := &Hash{}
	for i := 0; i <= 100000; i++ {
		h.Set(fmt.Sprintf(`%v`, i), i)
	}
}
