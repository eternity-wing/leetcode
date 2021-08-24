package treesandgraphs

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {
	numberOfItems := len(preorder)
	if numberOfItems == 0 {
		return nil
	}

	inorderIndexMap := getIndexedMap(inorder)

	var buildTreeInRange func(startIndex int, endIndex int) *TreeNode
	buildTreeInRange = func(startIndex int, endIndex int) *TreeNode {
		if startIndex > endIndex {
			return nil
		}

		rootValue := preorder[startIndex]
		root := &TreeNode{Val: rootValue}

		if endIndex > startIndex {
			rootIndex := inorderIndexMap[rootValue]
			i := startIndex + 1

			for ; i < endIndex && rootIndex >= inorderIndexMap[preorder[i]]; i++ {
			}
			if rootIndex >= inorderIndexMap[preorder[i]] {
				root.Left = buildTreeInRange(startIndex+1, i)
			}else{
				root.Left = buildTreeInRange(startIndex+1, i - 1)
				root.Right = buildTreeInRange(i, endIndex)
			}
		}
		return root
	}

	root := buildTreeInRange(0, numberOfItems-1)

	return root
}

func getIndexedMap(arr []int) map[int]int {
	result := make(map[int]int, len(arr))
	for index, val := range arr {
		result[val] = index
	}
	return result
}

func RunBuildTree() {
	fmt.Printf("\n\n")
	x := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	//x := buildTree([]int{1, 2}, []int{2, 1})
	fmt.Printf("%+v", x)
	return
	printInOrder(x)
}

func printInOrder(root *TreeNode)  {
	if root == nil {
		fmt.Printf(" null ")
		return
	}
	fmt.Printf("%d -> ", root.Val)
	printInOrder(root.Left)
	printInOrder(root.Right)
}