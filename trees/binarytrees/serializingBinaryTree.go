package binarytrees

import (
	"fmt"
	"strconv"
)

const NIL_NOTATION = "*"
const SERILIZER_SEPARATOR = ","

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return NIL_NOTATION
	}
	return fmt.Sprintf("%d%s%s%s%s", root.Val, SERILIZER_SEPARATOR, this.serialize(root.Left), SERILIZER_SEPARATOR, this.serialize(root.Right))
}


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	result, _ := this.deserializeBinaryTree(data)
	return result
}


// Deserializes your encoded data to tree.
func (this *Codec) deserializeBinaryTree(data string) (root *TreeNode, remainNodes string) {
	rootValue, remainNodes := this.PopValue(data)
	if data == "" || rootValue == NIL_NOTATION {
		return root, remainNodes
	}
	intVal, _ := strconv.Atoi(rootValue)
	leftNode, remainNodes := this.deserializeBinaryTree(remainNodes)
	rightNode, remainNodes := this.deserializeBinaryTree(remainNodes)
	root = &TreeNode{
		Val:   intVal,
		Left:  leftNode,
		Right: rightNode,
	}

	return root, remainNodes
}

func (this *Codec) PopValue(data string) (val string, remain string) {
	lengthOfData := len(data)
	for index, char := range data{
		if string(char) == SERILIZER_SEPARATOR{
			if lengthOfData > int(index+1) {
				return data[:index], data[index+1:]
			}
			return data[:index], ""
		}
	}
	return data, remain
}


func RunCodec()  {
	three := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	myTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: three,
	}
	codec := Constructor()
	seralized := codec.serialize(myTree)
	fmt.Printf("%+v\n", codec.deserialize(seralized))
}