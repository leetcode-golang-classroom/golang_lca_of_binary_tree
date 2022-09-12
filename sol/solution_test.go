package sol

import (
	"reflect"
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	temp := make([]*TreeNode, len(*input))
	for idx, val := range *input {
		num := 0
		if val != "null" {
			num, _ = strconv.Atoi(val)
		}
		if num != 0 {
			temp[idx] = &TreeNode{Val: num}
		}
	}
	for idx, val := range temp {
		if val != nil {
			if idx == 0 {
				tree = val
			}
			if 2*idx+1 < len(temp) {
				val.Left = temp[2*idx+1]
			}
			if 2*idx+2 < len(temp) {
				val.Right = temp[2*idx+2]
			}
		}
	}
	return tree
}

func BenchmarkTest(b *testing.B) {
	root := CreateBinaryTree(&[]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 4}
	for idx := 0; idx < b.N; idx++ {
		lowestCommonAncestor(root, p, q)
	}
}
func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8",
			args: args{root: CreateBinaryTree(&[]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"}),
				p: &TreeNode{Val: 2}, q: &TreeNode{Val: 8},
			},
			want: &TreeNode{Val: 6},
		},
		{
			name: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4",
			args: args{root: CreateBinaryTree(&[]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"}),
				p: &TreeNode{Val: 2}, q: &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 2},
		},
		{
			name: "root = [2,1], p = 2, q = 1",
			args: args{root: CreateBinaryTree(&[]string{"2", "1"}),
				p: &TreeNode{Val: 2}, q: &TreeNode{Val: 1},
			},
			want: &TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
