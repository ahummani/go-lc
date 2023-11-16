func sortList(head *ListNode) *ListNode {
  if head==nil || head.Next==nil{
    return head
  }
  prev,slow,fast:= &ListNode{},head,head
  for fast !=nil && fast.Next !=nil {
    prev=slow
    slow=slow.Next
    fast=fast.Next.Next
  }

  prev.Next=nil
  left:=sortList(head)
  right:=sortList(slow)

  return merge(left,right)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    if l1.Val < l2.Val {
        l1.Next = merge(l1.Next, l2)
        return l1
    }
    l2.Next = merge(l1, l2.Next)
    return l2
}