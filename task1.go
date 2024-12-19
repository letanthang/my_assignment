// Task 1 : detect sub route, Solution 2: start from the last element of list
// So we have to build a reverseHead
package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {

}

var (
	head *ListNode
	reversedHead *ListNode
	root *TreeNode
	matched = false
)  

func searchBackward(treeNode *TreeNode, parentTreeNodes []TreeNode) {
	// we return immediately if matched !!!!
	if treeNode == nil || matched {
		return
	}

	parentTreeNodes = append(parentTreeNodes, *treeNode)
	
	if treeNode.Value == reversedHead.Value {
		matched = checkMatchReverse(parentTreeNodes, reversedHead)
	}
	
	
	searchBackward(treeNode.Left, parentTreeNodes)
	searchBackward(treeNode.Right, parentTreeNodes)
}

func checkMatchReverse(parentTreeNodes []TreeNode,  head *ListNode) bool {
	// quite simple to implement
	// we use a loop to compare each element

	// pseudo code - just to pass the compiler type check
	return parentTreeNodes[0].Value == head.Value 
}
