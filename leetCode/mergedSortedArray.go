package main
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1, nums2)
	fmt.Printf(nums1)
}

func main(){
	nums1 := [1,2,3]
	m := 3
	nums2 := [2,5,6]
	n := 3
	merge(nums1, m, nums2, n)
}