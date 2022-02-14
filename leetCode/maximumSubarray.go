func maxSubArray(nums []int) int {
	maxSum := 0

	if len(nums) == 0 {
		return maxSum
	}
	maxSum = nums[0]

	arraySlice := make([]int, len(nums))

	for i := 0; i < len(nums); i++{
		if maxSum + nums[i]
	}

	return maxSum
}