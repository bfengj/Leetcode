package question19

type ListNode struct {
	Val  int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hhead := &ListNode{0,head}
	left,right := hhead,head
	for i:=0;i<n;i++ {
		right = right.Next
	}
	for ;right!=nil;left=left.Next {
		right = right.Next
	}
	left.Next = left.Next.Next
	return hhead.Next
}

/*func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left,right := head,head
	for i:=0;i<n;i++ {
		right = right.Next
	}
	if right==nil {
		return head.Next
	}
	for ;right.Next!=nil;left=left.Next {
		right = right.Next
	}
	left.Next = left.Next.Next
	return head
}*/