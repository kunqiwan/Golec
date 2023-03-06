package main

import "fmt"


func maxConsecutiveAnswers(answerKey string, k int) int {

	targetT := singleChar(answerKey,k,"T")
	targetF := singleChar(answerKey,k,"F")
	fmt.Println(targetT)
	fmt.Println(targetF)
	if targetT > targetF{
		return targetT
	} else{
		return targetF

	}
}

func singleChar(answerKey string, k int,target string)int{
	left :=0
	right :=0
	count :=0
	curr :=0
	max :=0
	for right<len(answerKey){
		if target == string(answerKey[right]){
			count++
		}
		right++
		for count > k {
			if(string(answerKey[left]) == target){
				count--
			}
			left++
		}
		curr = right - left
		if curr >max {
			max = curr
		}
	}
	return max
}