package main


func maxVowels(s string, k int) int {
	left:=0
	right:=0
	count :=0
	max :=0
	for right<len(s){
		curr := s[right]
		if isVewol(rune(curr)){
			count++
		}
		right++
		if right - left ==k{
			if count >max{
				max = count
			}
			if isVewol(rune(s[left])){
				count--
			}
			left++
		}
	}
	return max
}

func isVewol(target rune) bool{
	if target == 'a'|| target == 'e' || target == 'i' || target == 'o'|| target == 'u'{
		return true
	} else {
		return false
	}
}