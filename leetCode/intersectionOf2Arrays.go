func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func intersect(nums1 []int, nums2 []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}