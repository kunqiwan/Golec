package main

import "fmt"

func main() {
	nums1 :=[]int{1,1,3,2}
twoOutOfThree(nums1,nums1,nums1)
}
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	save :=map[int]struct{}{}
	res :=[]int{}
	for n:=0;n<len(nums1);n++{
		_,has :=save[nums1[n]]
		if has{
		} else{
			save[nums1[n]] = struct{}{}
		}
	}
	for m:=0;m<len(nums2);m++{
		_,has :=save[nums2[m]]
		if has{
			res = append(res,nums2[m])
		} else{
			
		}
	}
	fmt.Println(save)
	return res
}