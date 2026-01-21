package main

import "math"

/*
Дано два отсортированных массива nums1 и nums2 размера m и n соответственно, верните медиану двух отсортированных массивов.
Общая сложность времени выполнения должна составлять O(log (m+n)).

Example 1:
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Объяснение: объединенный массив = [1,2,3] и медиана равна 2.

Example 2:
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Объяснение: объединенный массив = [1,2,3,4] и медиана равна (2 + 3)/2 = 2,5.

Ограничения:
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2
	if len(A) > len(B) {
		A, B = B, A
	}
	total := len(A) + len(B)
	half := total / 2
	l, r := 0, len(A) - 1
	for {
		// разделение должно быть перед началом массива в случае отрицательных
		i := (l+r)>>1 //A half floor division
		j := half - (i + 1) - 1 //B half (del zeros)
		var Aleft, Aright, Bleft, Bright int
		if i >= 0 {
			Aleft = A[i]
		} else {
			Aleft = math.MinInt
		}
		if (i + 1) < len(A) {
			Aright = A[i+1]
		} else {
			Aright = math.MaxInt
		}
		if j >= 0 {
			Bleft = B[j]
		} else {
			Bleft = math.MinInt
		}
		if (j + 1) < len(B) {
			Bright = B[j+1]
		} else {
			Bright = math.MaxInt
		}
		if Aleft <= Bright && Bleft <= Aright {
			if total % 2 == 1 {
				return float64(min(Aright, Bright)) 
			}
			return float64(max(Aleft, Bleft) + min(Aright, Bright))/2
		} else if Aleft > Bright {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
