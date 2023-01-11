package main

import "fmt"

func main() {
	s := "baa"
	p := "aa"
	fmt.Println(findAnagrams(s,p))
}

func findAnagrams(s string, p string) []int {
	needs := make(map[string]int,len(p))
	windows := make(map[string]int,len(p))
	typenumber := 0
	for i:=0;i<len(p);i++{
		curr := string(p[i])
		_,ok := needs[curr]
		if ok {
			needs[curr] = needs[curr]+1
		} else{
			needs[curr] = 1
			typenumber++
		}
	}
	fmt.Println("needs:",needs)
	left :=0
	right :=0
	valid :=0
	save := make([]int,0,len(p))
	count :=0
	for right < len(s){
		target :=string(s[right])
		right++
		_,ok :=needs[target]
		if ok{
			windows[target]++
			if windows[target] ==needs[target]{
				valid++
				fmt.Println("left valid:",valid)
				fmt.Println("map:",windows)
			}
			fmt.Println("left :",left)
			fmt.Println("right :",right)
			 for valid == typenumber{
				 fmt.Println("跑了valid循环")
				if right - left == len(p){

					save =append(save,left)
					fmt.Println("count",count)
				}
				now := string(s[left])
				left++
				_,ok := needs[now]
				if ok{
					if windows[now] == needs[now]{
						valid--
					}
					windows[now] --
				}
			 }
		}

	}
	return save
}