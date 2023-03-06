package main


func repeatedNTimes(nums []int) int {
require := len(nums)/2
save :=make(map[int]int)
for n:=0;n<len(nums);n++{
save[nums[n]]++
}
	for k, v := range save {
		if v == require {
			key := k
			return key
		}
	}
	return 0
}
