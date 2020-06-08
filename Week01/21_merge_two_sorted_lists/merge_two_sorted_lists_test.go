package _21_merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"list0",
			args{
				l1: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				l2: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
