package question116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left.Next = root
		root.Right.Next = root
	}
	if root.Next != nil {
		if root.Next.Right != root {
			root.Next = root.Next.Right
		} else if root.Next.Next == nil {
			root.Next = nil
		} else {
			root.Next = root.Next.Next.Left
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
