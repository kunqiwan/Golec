package main

import "fmt"

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println(characterReplacement(s,k))
}
func characterReplacement(s string, k int) int {
//本题的窗口大小不能按照固定长度的思路，而是要让它一直满足k值可以填充整个窗口中的非最大字符值的条件
	right:=0
	left:=0
	max :=0
	save :=[26]int{}
	for right<len(s){
		//右滑中遇到的字母频次+1
		save[s[right]-'A']++
		if save[s[right]-'A']>max{
			max = save[s[right]-'A']
		}
		right++
		if right - left>max+k{
			save[s[left]-'A']--
			left++
		}
	}
	return right-left
}
