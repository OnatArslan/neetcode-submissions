func getConcatenation(nums []int) []int {
    ans := make([]int, 2 * len(nums))
	for i, _ := range nums {
		ans[i] = nums[i]
		ans[len(nums) + i] = nums[i]
	}
	return ans
}
