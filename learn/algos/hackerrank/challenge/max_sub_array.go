package challenge

import (
	"math"
)

/**
	We define subsequence as any subset of an array. We define a subarray as a contiguous subsequence in an array.

	Given an array, find the maximum possible sum among:
    * all nonempty subarrays.
    * all nonempty subsequences.

	https://www.hackerrank.com/challenges/maxsubarray/problem
*/
func MaxSubArray(input []int) (result []int) {
	return KadensAlgorithm(input)
}

/**
https://www.youtube.com/watch?v=86CQq3pKSUw
*/
func KadensAlgorithm(input []int) (result []int) {
	sum, global_sum, nonContigousSum := input[0], input[0], 0

	for i, value := range input {
		/*
			#Mistake 2 Check this Commit
			Didn't correctly find nonContigous Sum
			Include only positives values and handle all negative values case
			at end.
		*/
		if value > 0 {
			nonContigousSum += value
		}

		if i > 0 {
			if sum+value < value {
				/* Max Subarray starts here, drop previous max subarray */
				sum = value
			} else {
				/*
					This element is part of max subarray hence previous max
					subarray plus this element
				*/
				sum += value
			}

			/*
				Global Sum can be more than sum in between
				as its not max sum at this index, its max sum
				till now over any index.
			*/
			if global_sum < sum {
				global_sum = sum
			}
			//fmt.Println("Trace:", i, sum, global_sum)
		}
	}

	/*
		#Mistake 2 Fix
		If all values are negative then we would have not included anything
		in contigous sum and it must be equal to global sum.

		If there is even one positive value then global_sum can't be less than 0
	*/
	if global_sum < 0 {
		nonContigousSum = global_sum
	}

	return []int{global_sum, nonContigousSum}
}

/**
Brute Force O(n^2)
*/
func MaxSubArrayBruteForce(input []int) (result []int) {
	/*
		Mistake #1 as array can have negative numbers
		Sums should start negative.
	*/
	contigousSum := -math.MaxInt32
	nonContigousSum := 0
	if input[0] < 0 {
		nonContigousSum = input[0]
	}
	n := len(input)
	for i := 0; i < n; i++ {
		sum := 0
		/* Consider Segment from i to j over all possiblities */
		for j := i; j < n; j++ {
			sum += input[j]
			if sum > contigousSum {
				/* Subarry elements must be placed next to each other */
				contigousSum = sum
			}
			//fmt.Println(i, j, input[i:j+1], sum, contigousSum)
		}

		/* Sub Segment Sum may have gaps */
		if nonContigousSum < nonContigousSum+input[i] {
			nonContigousSum += input[i]
		}
	}
	return []int{contigousSum, nonContigousSum}
}
