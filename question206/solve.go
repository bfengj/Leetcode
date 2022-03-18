package question206

type ListNode struct {
	Val  int
	Next *ListNode
}
//迭代
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr!=nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}


//迭代
/*func reverseList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil {
		return head
	}
	var mid,right *ListNode
	mid = head.Next
	right = mid.Next
	head.Next = nil
	for right!=nil{
		mid.Next = head
		head = mid
		mid = right
		right = right.Next
	}
	mid.Next = head
	return mid
}*/

//递归
/*func reverseList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	//很重要，为了防止链表循环，在处理第一个节点的时候 防止循环
	head.Next = nil
	return res
}*/