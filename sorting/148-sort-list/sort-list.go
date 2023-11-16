func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    arr := make([]int, 0)

    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    //sort it 

    sort.Ints(arr)

    head = nil

    for i:= len(arr)-1; i>=0; i-- {
        if i == len(arr)-1 { 
            head = nil
        }
        newNode := &ListNode{Val:arr[i]}
        newNode.Next = head
        head = newNode
    }
    return head
}