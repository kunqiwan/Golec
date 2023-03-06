package main


func findDisappearedNumbers(nums []int) []int {
right :=0
result :=[]int{}
for right<len(nums){
curr := nums[right]
if curr == right+1{
	right++
} else if curr !=right+1{
	temp :=nums[right]
	if nums[right] == nums[temp-1]{
		//说明这位置上有已存在值，多余了
		right++
	} else {
		nums[right] = nums[temp-1]
		nums[temp-1] = temp
	}
	//fmt.Println("跑了一次",nums)
}

}
	for n:=0;n<len(nums);n++{
		if nums[n]!=n+1{
			result =append(result,n+1)
		}
	}
	return result

}
