func isOdd(n int) bool {
	return n%2 != 0
}

func median(A []int, l, r int) float64 {
	length := r - l + 1
	m2 := (r + l + 1) / 2
	m1 := m2 - 1

	if isOdd(length) {
		return float64(A[m2])
	}

	return float64(A[m1]+A[m2]) / 2.0
}

func binarySearch(A []int, l, r, val int, left bool) int {
	for l <= r {
		i := (r + l + 1) / 2

		if A[i] == val {

			if left {
				for j := i + 1; j <= r; j++ {
					if A[j] != val {
						break
					}

					i = j
				}

				return i + 1
			}

			for j := i - 1; j >= l; j-- {
				if A[j] != val {
					break
				}

				i = j
			}

			return i
		}

		if A[i] > val {
			r = i - 1
			continue
		}

		l = i + 1
	}

	return l
}

// i -> index of value "v" in its own array
// A -> foreign array
// n -> length of the merged array from which we want to figure the median out
// left -> search to the left of the median
// @return -> if the real index of "v" is behind or ahead of the median
func isValueAheadOfMedian(A []int, l, r, i, v, n int, left bool) bool {
	median := n / 2

	iA := binarySearch(A, l, r, v, left)
	i += iA // real index of "v"

	return i >= median
}

func findMedianSubsequences(A, B []int, lA, rA, lB, rB int) float64 {
	m := len(A)
	n := len(B)
	o := m + n
	mi1 := o / 2
	mi2 := o/2 - 1

	var med1, med2 int

	// prevent searching if 2 medians have already been determined
	flag1 := false
	flag2 := false

	for i := lA; i < rA+1; i++ {
		if flag1 && flag2 {
			break
		}

		// find the real index of this value
		v := A[i]
		iA := binarySearch(B, lB, rB, v, true)
		ri := i + iA

		if ri == mi1 {
			if isOdd(o) {
				return float64(v)
			}

			med1 = v
			flag1 = true
			continue
		}

		if ri == mi2 {
			med2 = v
			flag2 = true
		}
	}

	for i := lB; i < rB+1; i++ {
		if flag1 && flag2 {
			break
		}

		v := B[i]
		iB := binarySearch(A, lA, rA, v, false)
		ri := i + iB

		if ri == mi1 {
			if isOdd(o) {
				return float64(v)
			}

			med1 = v
			flag1 = true
			continue
		}

		if ri == mi2 {
			med2 = v
			flag2 = true
		}
	}

	if isOdd(o) || !flag1 || !flag2 {
		panic("Couldn't find median on final subsequence...")
	}

	return float64(med1+med2) / 2.0
}

func findMedianSortedArrays(A []int, B []int) float64 {
	m := len(A)
	n := len(B)
	o := m + n

	if len(A) <= 0 {
		return median(B, 0, len(B)-1)
	}

	if len(B) <= 0 {
		return median(A, 0, len(A)-1)
	}

	targetLen := 2

	lA := 0
	rA := len(A) - 1
	lB := 0
	rB := len(B) - 1

	for {
		lenA := rA - lA + 1
		lenB := rB - lB + 1

		if lenA+lenB <= targetLen*2 {
			// Get the real index of each value of the final subsequence.
			// The median will be included in this subsequence.
			return findMedianSubsequences(A, B, lA, rA, lB, rB)
		}

		// Cut both subsequences in half until they have a small constant length.
		// Always choose the half of the sequence which includes the global median.
		//
		// The real index of a value is its index in the merged array of A and B.
		// To find the real index of a value in sequence A perform binarySearch
		//  on B and add both indices.
		//
		// If the real index is lower than the median's index, the median must be
		//  somewhere in the upper half of the sequence.

		if lenA > targetLen {
			m := (rA + lA + 1) / 2

			if isValueAheadOfMedian(B, lB, rB, m, A[m], o, true) {
				rA = m
			} else {
				lA = m
			}
		}

		if lenB > targetLen {
			m := (rB + lB + 1) / 2

			if isValueAheadOfMedian(A, lA, rA, m, B[m], o, false) {
				rB = m
			} else {
				lB = m
			}
		}
	}
}