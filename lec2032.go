package main

func main() {
	nums1 :=[]int{1,1,3,2}
twoOutOfThree(nums1,nums1,nums1)
}
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
s1 :=exist(nums1)
s2 :=exist(nums2)
s3 :=exist(nums3)
res := []int{}
for m:=0;m<len(s1);m++{
	if s1[m]+s2[m]+s3[m]>1{
		res = append(res,m)
	}
}
return res
}
func exist(nums []int)[101]int {
	save :=[101]int{}
for n:=0;n<len(nums);n++{
	save[nums[n]] = 1
}
return save
}