package main


func smallerNumbersThanCurrent(nums []int) []int {
frequ :=[101]int{}
for n:=0;n<len(nums);n++{
	//记录每个数出现的频次
frequ[nums[n]]++
}
//记录每个数前面存在的数比它小的有多少
for i:=1;i<len(frequ);i++{
	frequ[i]=frequ[i]+frequ[i-1]
}
ans := make([]int, len(nums))
for m:=0;m<len(ans);m++{
	curr:=nums[m]
	if curr>0{
		ans[m] = frequ[curr-1]
	}
}
return  ans
}
