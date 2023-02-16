package main

import "fmt"

func main() {
	nums := []int{4,2,4,5,6}
	fmt.Println(maximumUniqueSubarray(nums))
}
func maximumUniqueSubarray(nums []int) int {
left:=0
right:=0
save :=make(map[int]int,len(nums))
sum :=0
max := -1
for right<len(nums){
	curr := nums[right]
	save[curr]++
	sum = sum+curr
	fmt.Println("curr is",curr)
	fmt.Println("sum is",sum)

	fmt.Println("save is",save)
	right++
	//挪动左窗口指针
	for save[curr] >1 {
		fmt.Println("left is",left)
		save[nums[left]]--
		sum = sum- nums[left]
		left++
	}

if max<sum {
	max = sum
}
	fmt.Println("max is",max)

}
return max
}

