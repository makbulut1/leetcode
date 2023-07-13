/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    newNode := &ListNode{}
    tmp := newNode
    tmp = newNode
    for  l1 != nil || l2 != nil{
        //l1.Val + l2.Val + extra = total
        if l1 != nil {
            newNode.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            newNode.Val += l2.Val
            l2 = l2.Next
        }
        // total check one step if not one step total 
        if newNode.Val > 9 {
            newNode.Val -= 10
            newNode.Next = &ListNode{Val:1}
        } else if l1 != nil || l2 != nil{
            newNode.Next = &ListNode{}
        }
        //next after node
        newNode = newNode.Next
    }
    newNode = tmp
    return newNode
}
