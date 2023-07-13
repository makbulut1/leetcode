# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
The problem requires adding two numbers represented as linked lists. Each node in the linked list represents a digit, and the digits are stored in reverse order. For example, the number 123 is represented as 3->2->1.

# Approach
<!-- Describe your approach to solving the problem. -->
The approach to solving this problem involves iterating over the linked lists simultaneously, adding the corresponding digits and maintaining a carry for the next sum. The sum of two digits (l1.Val and l2.Val) plus the carry will be the value of the new node. If the sum exceeds 9, we need to adjust the value and carry accordingly.

The steps for the approach are as follows:

1. Create a new node to store the result.
2. Initialize a temporary variable, tmp, to keep track of the new node.
3. Iterate over the linked lists while at least one of them is not empty:
    - Add the values of the current nodes from both linked lists to the value of the new node.
    - Move to the next node in each linked list if it exists.
    - Check if the sum of the values plus any carry is greater than 9. If so, subtract 10 from the sum and set the next node as a new node with a value of 1 (carry).
    - If either linked list has more nodes, create a new node for the next digit.
    - Move to the next node in the new linked list.
Return the new linked list as the sum of the two numbers.

# Complexity
- Time complexity: The algorithm iterates through the linked lists once, which takes O(max(m, n)) time, where m and n are the lengths of the input linked lists l1 and l2, respectively.
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: The algorithm creates a new linked list to store the sum, which takes O(max(m, n)) space in the worst case if there is a carry to the next digit.
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
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
```
