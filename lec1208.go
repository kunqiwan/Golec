package main

import "math"


func equalSubstring(s string, t string, maxCost int) int {
	left:=0
	right:=0
	count :=0
	max :=0
	for right<len(s){
		if s[right] != t[right]{
			a :=math.Abs(float64(rune(s[right]) - rune(t[right])))
			count = count+int(a)
		}
		right++
		for count > maxCost{
			if s[left] !=t[left]{
				b :=math.Abs(float64(rune(s[left]) - rune(t[left])))
				count = count-int(b)
			}
			left++
		}
		if right-left >max{
			max = right -left
		}

	}
	return max
}