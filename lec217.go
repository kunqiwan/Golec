package main



//func containsDuplicate(nums []int) bool {
//save :=make(map[int]int)
//for n:=0;n<len(nums);n++{
//	save[nums[n]]++
//	if save[nums[n]]>1{
//		return true
//	}
//}
//return  false
//}

func containsDuplicate(nums []int) bool {
	save :=map[int]struct{}{}
	for n:=0;n<len(nums);n++{
		_,has :=save[nums[n]]
		if has{
			return true
		}
		save[nums[n]] = struct{}{}
	}

return false
}