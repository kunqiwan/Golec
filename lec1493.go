package main


func longestSubarray(nums []int) int {
left :=0
right :=0
count :=0
max := 0
for right<len(nums){
if nums[right] ==0{
	count++
}
right++
for count >1{
if nums[left] == 0{
	count--
}
left++
}
curr_len := right - left
if max <curr_len{
	max = curr_len
}

}
if count ==0{
	return len(nums)-1
}
return max-1
}
