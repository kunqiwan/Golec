package main

func main() {

}
func longestOnes(nums []int, k int) int {
left:=0
right:=0
count :=0
max :=0
for right<len(nums){
	if nums[right] == 0{
		count++
	}
	right++
	for count >k{
		if nums[left] ==0{
			count--
		}
		left++
	}
	if right-left > max{
		max = right-left
	}
}
return  max
}
