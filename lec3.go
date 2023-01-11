package main

import (
	"fmt"
)

func main() {
	s := "abcsdsad"
	fmt.Println(s)
}
func lengthOfLongestSubstring(s string) int {
windows := make(map[string]int,len(s))
left :=0
right :=0
maxcount := 0
for right < len(s){
	 curr :=string(s[right])
	 fmt.Println("right:",right)
	 right++
	 _,ok:=windows[curr]
	 if ok{
		 for left < right-1{
			 currlen := right-left-1
			 if currlen>maxcount{
				 fmt.Println(windows)
				 maxcount = currlen
			 }
			 leftmove := string(s[left])
			 left++
			 windows[leftmove] =windows[leftmove]-1
			 fmt.Println("left",left)
		 }
	 } else {
		windows[curr] = windows[curr]+1
	 }
}
return  maxcount
}
