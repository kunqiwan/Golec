package main

import "fmt"


func numOfSubarrays(arr []int, k int, threshold int) int {
	left :=0
	right :=0
	sum := 0
	count :=0
	for right<len(arr){
		sum = sum+arr[right]
		right++
		if right -left ==k{
			av :=sum/k
			if(av>=threshold){
				fmt.Println("左：",left)
				fmt.Println("右:",right)
				count++
			}
			sum = sum-arr[left]
			left++
		}
	}
	return count
}
