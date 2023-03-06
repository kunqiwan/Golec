package main


//func majorityElement(nums []int) int {
//length := len(nums)
//req := length/2
//save := make(map[int]int,len(nums))
//for n:=0;n<length;n++{
//	save[nums[n]]++
//}
//	for key, element := range save {
//		if element>req{
//			return key
//		}
//	}
//	return 0
//}

func majorityElement(nums []int) int {
	major :=0
	count :=0
	for n:=0;n<len(nums);n++{
		curr :=nums[n]
		if count ==0{
			major =curr
		}
		if curr ==major{
			count++
		} else{
			count--
		}
	}
	return major
}