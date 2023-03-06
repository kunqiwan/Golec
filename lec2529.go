package main

//For Loop Method

//func maximumCount(nums []int) int {
//	neg:=0
//	pos:=0
//for n:=0;n<len(nums);n++{
//	if nums[n] <0{
//		neg++
//	}
//	if nums[n]>0{
//		pos++
//	}
//}
//if neg>=pos{
//	return neg
//} else{
//	return pos
//}
//}

func maximumCount(nums []int) int {
	nr:=neg(nums)
	pr :=pos(nums)
	if nr>=pr{
		return nr
	} else {
		return  pr
	}
}
func neg(nums []int) int {
	length :=len(nums)
	right := length-1
	left :=0
	//find negative range,right range
	for left<=right{
		mid :=left+(right-left)/2
		if nums[mid]>0{
			right = mid-1
		} else if nums[mid] ==0{
			right = mid -1
		} else if nums[mid] <0{
			left =mid+1
		}
	}
	nr := right+1
	return nr
}
func pos(nums []int)int{
	length :=len(nums)
	right := length-1
	left :=0
	//find positive range， left range
	for left<=right{
		mid :=left+(right-left)/2
		if nums[mid]<=0{
			left = mid+1

		} else {
			right = mid-1
		}
	}
	pr:=length-left
	return pr

}