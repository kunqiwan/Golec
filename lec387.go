package main

import "fmt"

func firstUniqChar(s string) int {
first := [26]int{}
one := s[0]
for n:=1;n<len(s);n++{
	curr :=rune(s[n])
	if first[curr-'a']==0{
		first[curr-'a'] = n
	} else if first[curr-'a']<=1000 && first[curr-'a']!=0{
		first[curr-'a'] =100050
	}
	if rune(one) == curr{
		first[curr-'a'] =100050
	}

	}
	if first[rune(one)-'a'] ==0{
		return 0
	}
max :=100005
for m:=0;m<len(first);m++{
	//fmt.Println("first[m]",first[m])
	if first[m]<=100010 && first[m]!=0{
		if first[m]<max{
			max = first[m]
		}
		//fmt.Println("max",max)
	}
}
fmt.Println("Final Slice:",first)
if max==100005{
	return -1
} else {
	return max
}
}
