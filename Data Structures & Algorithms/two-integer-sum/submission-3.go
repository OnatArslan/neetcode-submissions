func twoSum(nums []int, target int) []int {
    mp := make(map[int]int, len(nums))

	for i,num := range nums {
		rem := target - num
		if j, ok := mp[rem];ok{
			return []int{j,i}
		}
		mp[num] = i
	}
	return []int{}
}
