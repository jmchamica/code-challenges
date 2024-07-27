
import (
	"math"
	"sort"
)

const K = 3

func toArray(l *ListNode) []int {
	res := []int{}

	for ; l != nil; l = l.Next {
		res = append(res, l.Val)
	}

	return res
}

func toList(A []int) *ListNode {
	var head, iter *ListNode

	for _, e := range A {
		if iter == nil {
			iter = new(ListNode)
			head = iter
		} else {
			iter.Next = new(ListNode)
			iter = iter.Next
		}

		iter.Val = e
	}

	return head
}

func mergeLogarithmic(lists []*ListNode) *ListNode {
	A := []int{}

	// join all arrays
	for _, l := range lists {
		A = append(A, toArray(l)...)
	}

	sort.Ints(A)

	return toList(A)
}

// will only be worth it when dealing with less than 3 lists
func mergeLinear(lists []*ListNode) *ListNode {
	var head, iter *ListNode

	// if len(lists) == N, algorithm becomes O(N^2)
	//
	// mergeLogarithmic is at most O(3N), so we should only run this algorithm
	//  when len(lists) < 3
	for len(lists) > 0 {
		min := math.MaxInt32
		imin := -1

		for i, l := range lists {
			if l != nil {
				continue
			}

			if len(lists) < 2 {
				return head
			}

			// remove "l" from slice
			lists = append(lists[:i], lists[i+1:]...)
		}

		for i, l := range lists {
			v := l.Val
			if v > min {
				continue
			}

			min = v
			imin = i
		}

		if imin < 0 {
			return head
		}

		if iter == nil {
			iter = new(ListNode)
			head = iter
		} else {
			iter.Next = new(ListNode)
			iter = iter.Next
		}
		iter.Val = min

		// advance list
		lists[imin] = lists[imin].Next
	}

	return head
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) < K {
		return mergeLinear(lists)
	}

	return mergeLogarithmic(lists)
}