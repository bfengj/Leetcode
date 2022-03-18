package question876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var leftNode,rightNode *ListNode = head,head
	for  {
		if rightNode.Next==nil {
			return leftNode
		}else{
			rightNode = rightNode.Next.Next
			leftNode = leftNode.Next
		}
		if rightNode==nil {
			return leftNode
		}
	}
}
