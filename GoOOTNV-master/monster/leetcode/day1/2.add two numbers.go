/**
 * @Difficulty: Medium
 * @Description: You are given two non-empty linked lists representing two non-negative integers.
 *			    The digits are stored in reverse order and each of their nodes contain a single digit.
 *			    Add the two numbers and return it as a linked list.
 *			    You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *				给出两个非空链表, 表示两个非负整数。数字以相反的顺序存储, 每个节点都包含一个数字。添加两个数字并将其作为链接列表返回。
 *				您可以假定两个数字不包含任何前导零, 除了数字0本身。
 * @Example: Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 *           Output: 7 -> 0 -> 8
 *           Explanation: 342 + 465 = 807.
 */


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	result := &ListNode{}
	result.Val = l1.Val + l2.Val

	if result.Val > 9 {
		result.Val -= 10

		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l2.Next.Val += 1
	}

	if l2.Next != nil || l1.Next != nil {
		result.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return result
}
