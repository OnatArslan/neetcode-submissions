func majorityElement(nums []int) int {
    moreThan := float64(len(nums)) / 2

	mp := make(map[int]int, len(nums))

	for _, num := range nums {
		mp[num]++
		if float64(mp[num]) > moreThan{
			return num
		}
	}
return 0
}
