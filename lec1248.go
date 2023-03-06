package main


func numberOfSubarrays(nums []int, k int) int {
	left :=0
	right:=0
	odd_count :=0
	result :=0

	for right<len(nums){
		if nums[right] %2 !=0{
			odd_count++
		}
		right++
		if odd_count ==k{
			right_init := right
			for right<len(nums)&&nums[right]%2==0  {
				right++
			}
			//在K限度内，right能达到的最大长度
			right_max := right-right_init
			left_init :=0
			for nums[left] %2 ==0{
				left_init ++
				left++
			}

			result = result+(right_max+1)*(left_init+1)
			left++
			odd_count --
		}

		// 先判断是否到达末尾

	}
	return result
}


