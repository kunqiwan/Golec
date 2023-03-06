package main


func intersection(nums1 []int, nums2 []int) []int {
save :=make(map[int]int,len(nums1))
check :=make(map[int]int,len(nums1))
result := []int{}
for n:=0;n<len(nums1);n++{
save[nums1[n]]++
}
for m:=0;m<len(nums2);m++{
	curr :=nums2[m]
	if save[nums2[m]] >0 && check[curr]==0{
		check[curr]++
		result = append(result,curr )
	}
}
return result
}
